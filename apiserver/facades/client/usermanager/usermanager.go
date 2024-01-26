// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package usermanager

import (
	"context"
	"time"

	"github.com/juju/errors"
	"github.com/juju/loggo/v2"
	"github.com/juju/names/v5"

	"github.com/juju/juju/apiserver/authentication"
	"github.com/juju/juju/apiserver/common"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/core/permission"
	coreuser "github.com/juju/juju/core/user"
	"github.com/juju/juju/internal/auth"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/state"
)

// UserService defines the methods to operate with the database.
type UserService interface {
	GetAllUsers(ctx context.Context) ([]coreuser.User, error)
	GetUserByName(ctx context.Context, name string) (coreuser.User, error)
	AddUserWithPassword(ctx context.Context, username, displayName string, createdBy coreuser.UUID, password auth.Password) (coreuser.UUID, error)
	AddUserWithActivationKey(ctx context.Context, name string, displayName string, creatorUUID coreuser.UUID) ([]byte, coreuser.UUID, error)
	EnableUserAuthentication(ctx context.Context, uuid coreuser.UUID) error
	DisableUserAuthentication(ctx context.Context, uuid coreuser.UUID) error
	SetPassword(ctx context.Context, uuid coreuser.UUID, password auth.Password) error
	ResetPassword(ctx context.Context, uuid coreuser.UUID) ([]byte, error)
	RemoveUser(ctx context.Context, uuid coreuser.UUID) error
}

// UserManagerAPI implements the user manager interface and is the concrete
// implementation of the api end point.
type UserManagerAPI struct {
	state       *state.State
	userService UserService
	pool        *state.StatePool
	authorizer  facade.Authorizer
	check       *common.BlockChecker
	apiUser     names.UserTag
	isAdmin     bool
	logger      loggo.Logger
}

// NewAPI creates a new API endpoint for calling user manager functions.
func NewAPI(
	state *state.State,
	userService UserService,
	pool *state.StatePool,
	authorizer facade.Authorizer,
	check *common.BlockChecker,
	apiUser names.UserTag,
	isAdmin bool,
	logger loggo.Logger,
) (*UserManagerAPI, error) {
	return &UserManagerAPI{
		state:       state,
		userService: userService,
		pool:        pool,
		authorizer:  authorizer,
		check:       check,
		apiUser:     apiUser,
		isAdmin:     isAdmin,
		logger:      logger,
	}, nil
}

func (api *UserManagerAPI) hasControllerAdminAccess() (bool, error) {
	err := api.authorizer.HasPermission(permission.SuperuserAccess, api.state.ControllerTag())
	return err == nil, err
}

// AddUser adds a user with a username, and either a password or
// a randomly generated secret key which will be returned.
func (api *UserManagerAPI) AddUser(ctx context.Context, args params.AddUsers) (params.AddUserResults, error) {
	var result params.AddUserResults

	if err := api.check.ChangeAllowed(ctx); err != nil {
		return result, errors.Trace(err)
	}

	if len(args.Users) == 0 {
		return result, nil
	}

	// Create the results list to populate.
	result.Results = make([]params.AddUserResult, len(args.Users))

	if _, err := api.hasControllerAdminAccess(); err != nil {
		return result, err
	}

	for i, arg := range args.Users {
		// Get creator user
		creator, err := api.userService.GetUserByName(ctx, api.apiUser.Id())
		if err != nil {
			err = errors.Annotate(err, "failed to get creator user")
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}

		// Add new user
		var newUserUUID coreuser.UUID
		var activationKey []byte
		if arg.Password != "" {
			newUserUUID, err = api.userService.AddUserWithPassword(ctx, arg.Username, arg.DisplayName, creator.UUID, auth.NewPassword(arg.Password))
		} else {
			activationKey, newUserUUID, err = api.userService.AddUserWithActivationKey(ctx, arg.Username, arg.DisplayName, creator.UUID)
		}
		if err != nil {
			err = errors.Annotate(err, "failed to create user")
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		} else {
			result.Results[i] = params.AddUserResult{
				Tag:       coreuser.UserTag(newUserUUID).String(),
				SecretKey: activationKey,
			}
		}
	}
	return result, nil
}

// RemoveUser permanently removes a user from the current controller for each
// entity provided. While the user is permanently removed we keep it's
// information around for auditing purposes.
// TODO(redir): Add information about getting deleted user information when we
// add that capability.
func (api *UserManagerAPI) RemoveUser(ctx context.Context, entities params.Entities) (params.ErrorResults, error) {
	var deletions params.ErrorResults

	if err := api.check.ChangeAllowed(ctx); err != nil {
		return deletions, errors.Trace(err)
	}

	controllerOwner, err := api.state.ControllerOwner()
	if err != nil {
		return deletions, errors.Trace(err)
	}

	// Create the results list to populate.
	deletions.Results = make([]params.ErrorResult, len(entities.Entities))

	isSuperUser, err := api.hasControllerAdminAccess()
	if err != nil {
		return deletions, errors.Trace(err)
	}
	if !api.isAdmin && !isSuperUser {
		return deletions, apiservererrors.ErrPerm
	}

	// Remove the entities.
	for i, e := range entities.Entities {
		user, err := names.ParseUserTag(e.Tag)
		if err != nil {
			deletions.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}

		if controllerOwner.Id() == user.Id() {
			deletions.Results[i].Error = apiservererrors.ServerError(
				errors.Errorf("cannot delete controller owner %q", user.Name()))
			continue
		}

		// Get User
		usr, err := api.userService.GetUserByName(ctx, user.Name())
		if err != nil {
			deletions.Results[i].Error = apiservererrors.ServerError(
				errors.Annotatef(err, "failed to delete user %q", user.Name()))
			continue
		}

		// Remove user
		err = api.userService.RemoveUser(ctx, usr.UUID)
		if err != nil {
			deletions.Results[i].Error = apiservererrors.ServerError(
				errors.Annotatef(err, "failed to delete user %q", user.Name()))
			continue
		}
		deletions.Results[i].Error = nil

	}
	return deletions, nil
}

func (api *UserManagerAPI) getUser(tag string) (*state.User, error) {
	userTag, err := names.ParseUserTag(tag)
	if err != nil {
		return nil, errors.Trace(err)
	}

	// TODO(anvial): Legacy block to delete when user domain wire up is complete.
	user, err := api.state.User(userTag)
	if err != nil {
		return nil, errors.Wrap(err, apiservererrors.ErrPerm)
	}
	// End legacy block.

	return user, nil
}

// EnableUser enables one or more users.  If the user is already enabled,
// the action is considered a success.
func (api *UserManagerAPI) EnableUser(ctx context.Context, users params.Entities) (params.ErrorResults, error) {
	if _, err := api.hasControllerAdminAccess(); err != nil {
		return params.ErrorResults{}, err
	}

	if err := api.check.ChangeAllowed(ctx); err != nil {
		return params.ErrorResults{}, errors.Trace(err)
	}
	return api.enableUserImpl(ctx, users, "enable", (*state.User).Enable)
}

// DisableUser disables one or more users.  If the user is already disabled,
// the action is considered a success.
func (api *UserManagerAPI) DisableUser(ctx context.Context, users params.Entities) (params.ErrorResults, error) {
	if _, err := api.hasControllerAdminAccess(); err != nil {
		return params.ErrorResults{}, err
	}

	if err := api.check.ChangeAllowed(ctx); err != nil {
		return params.ErrorResults{}, errors.Trace(err)
	}
	return api.enableUserImpl(ctx, users, "disable", (*state.User).Disable)
}

func (api *UserManagerAPI) enableUserImpl(ctx context.Context, args params.Entities, action string, method func(*state.User) error) (params.ErrorResults, error) {
	var result params.ErrorResults

	if len(args.Entities) == 0 {
		return result, nil
	}

	if !api.isAdmin {
		if _, err := api.hasControllerAdminAccess(); err != nil {
			return result, err
		}
	}

	// Create the results list to populate.
	result.Results = make([]params.ErrorResult, len(args.Entities))

	for i, arg := range args.Entities {
		// Get User
		usr, err := api.userService.GetUserByName(ctx, arg.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(errors.Errorf("failed to %s user: %s", action, err))
			continue
		}

		// Enable/Disable user
		if action == "enable" {
			err = api.userService.EnableUserAuthentication(ctx, usr.UUID)
		} else {
			err = api.userService.DisableUserAuthentication(ctx, usr.UUID)
		}
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(errors.Errorf("failed to %s user: %s", action, err))
			continue
		}
	}
	return result, nil
}

// UserInfo returns information on a user.
func (api *UserManagerAPI) UserInfo(ctx context.Context, request params.UserInfoRequest) (params.UserInfoResults, error) {
	var results params.UserInfoResults
	isAdmin, err := api.hasControllerAdminAccess()
	if err != nil && !errors.Is(err, authentication.ErrorEntityMissingPermission) {
		return results, errors.Trace(err)
	}

	var accessForUser = func(userTag names.UserTag, result *params.UserInfoResult) {
		// Lookup the access the specified user has to the controller.
		access, err := common.GetPermission(api.state.UserPermission, userTag, api.state.ControllerTag())
		if err == nil {
			result.Result.Access = string(access)
		} else if err != nil && !errors.Is(err, errors.NotFound) {
			result.Result = nil
			result.Error = apiservererrors.ServerError(err)
		}
	}

	// TODO(anvial): Legacy block to delete when user domain wire up is complete.
	var infoForUser = func(user *state.User) params.UserInfoResult {
		var lastLogin *time.Time
		userLastLogin, err := user.LastLogin()
		if err != nil {
			if !state.IsNeverLoggedInError(err) {
				api.logger.Debugf("error getting last login: %v", err)
			}
		} else {
			lastLogin = &userLastLogin
		}
		result := params.UserInfoResult{
			Result: &params.UserInfo{
				Username:       user.Name(),
				DisplayName:    user.DisplayName(),
				CreatedBy:      user.CreatedBy(),
				DateCreated:    user.DateCreated(),
				LastConnection: lastLogin,
				Disabled:       user.IsDisabled(),
			},
		}
		if user.IsDisabled() {
			// disabled users have no access to the controller.
			result.Result.Access = string(permission.NoAccess)
		} else {
			accessForUser(user.UserTag(), &result)
		}
		return result
	}
	// End legacy block.

	argCount := len(request.Entities)
	if argCount == 0 {

		// TODO(anvial): Legacy block to delete when user domain wire up is complete.
		users, err := api.state.AllUsers(request.IncludeDisabled)
		if err != nil {
			return results, errors.Trace(err)
		}
		for _, user := range users {
			if !isAdmin && !api.authorizer.AuthOwner(user.Tag()) {
				continue
			}
			results.Results = append(results.Results, infoForUser(user))
		}
		// End legacy block.

		return results, nil
	}

	// Create the results list to populate.
	for _, arg := range request.Entities {
		userTag, err := names.ParseUserTag(arg.Tag)
		if err != nil {
			results.Results = append(results.Results, params.UserInfoResult{Error: apiservererrors.ServerError(err)})
			continue
		}
		if !isAdmin && !api.authorizer.AuthOwner(userTag) {
			results.Results = append(results.Results, params.UserInfoResult{Error: apiservererrors.ServerError(apiservererrors.ErrPerm)})
			continue
		}
		if !userTag.IsLocal() {
			// TODO(wallyworld) record login information about external users.
			result := params.UserInfoResult{
				Result: &params.UserInfo{
					Username: userTag.Id(),
				},
			}
			accessForUser(userTag, &result)
			results.Results = append(results.Results, result)
			continue
		}

		// TODO(anvial): Legacy block to delete when user domain wire up is complete.
		user, err := api.getUser(arg.Tag)
		if err != nil {
			results.Results = append(results.Results, params.UserInfoResult{Error: apiservererrors.ServerError(err)})
			continue
		}
		results.Results = append(results.Results, infoForUser(user))
		// End legacy block.
	}

	return results, nil
}

func (api *UserManagerAPI) checkCanRead(modelTag names.Tag) error {
	return api.authorizer.HasPermission(permission.ReadAccess, modelTag)
}

// ModelUserInfo returns information on all users in the model.
func (api *UserManagerAPI) ModelUserInfo(ctx context.Context, args params.Entities) (params.ModelUserInfoResults, error) {
	var result params.ModelUserInfoResults
	for _, entity := range args.Entities {
		modelTag, err := names.ParseModelTag(entity.Tag)
		if err != nil {
			return result, errors.Trace(err)
		}
		infos, err := api.modelUserInfo(modelTag)
		if err != nil {
			return result, errors.Trace(err)
		}
		result.Results = append(result.Results, infos...)
	}
	return result, nil
}

func (api *UserManagerAPI) modelUserInfo(modelTag names.ModelTag) ([]params.ModelUserInfoResult, error) {
	var results []params.ModelUserInfoResult
	model, closer, err := api.pool.GetModel(modelTag.Id())
	if err != nil {
		return results, errors.Trace(err)
	}
	defer closer.Release()
	if err := api.checkCanRead(model.ModelTag()); err != nil {
		return results, err
	}

	users, err := model.Users()
	if err != nil {
		return results, errors.Trace(err)
	}

	for _, user := range users {
		var result params.ModelUserInfoResult
		userInfo, err := common.ModelUserInfo(user, model)
		if err != nil {
			result.Error = apiservererrors.ServerError(err)
		} else {
			userInfo.ModelTag = modelTag.String()
			result.Result = &userInfo
		}
		results = append(results, result)
	}
	return results, nil
}

// SetPassword changes the stored password for the specified users.
func (api *UserManagerAPI) SetPassword(ctx context.Context, args params.EntityPasswords) (params.ErrorResults, error) {
	if err := api.check.ChangeAllowed(ctx); err != nil {
		return params.ErrorResults{}, errors.Trace(err)
	}

	var result params.ErrorResults

	if len(args.Changes) == 0 {
		return result, nil
	}

	// Create the results list to populate.
	result.Results = make([]params.ErrorResult, len(args.Changes))
	for i, arg := range args.Changes {
		if err := api.setPassword(arg); err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
		}
	}
	return result, nil
}

func (api *UserManagerAPI) setPassword(arg params.EntityPassword) error {
	// TODO(anvial): Legacy block to delete when user domain wire up is complete.
	user, err := api.getUser(arg.Tag)
	if err != nil {
		return errors.Trace(err)
	}
	// End legacy block.

	if !api.isAdmin {
		if _, err := api.hasControllerAdminAccess(); err != nil && api.apiUser != user.UserTag() {
			return err
		}
	}

	if arg.Password == "" {
		return errors.New("cannot use an empty password")
	}

	// TODO(anvial): Legacy block to delete when user domain wire up is complete.
	if err := user.SetPassword(arg.Password); err != nil {
		return errors.Annotate(err, "failed to set password")
	}
	// End legacy block.

	// Get User
	usr, err := api.userService.GetUserByName(context.Background(), arg.Tag)
	if err != nil {
		return errors.Annotate(err, "failed to set password")
	}

	// Set password
	if err := api.userService.SetPassword(context.Background(), usr.UUID, auth.NewPassword(arg.Password)); err != nil {
		return errors.Annotate(err, "failed to set password")
	}

	return nil
}

// ResetPassword resets password for supplied users by
// invalidating current passwords (if any) and generating
// new random secret keys which will be returned.
// Users cannot reset their own password.
func (api *UserManagerAPI) ResetPassword(ctx context.Context, args params.Entities) (params.AddUserResults, error) {
	var result params.AddUserResults

	if err := api.check.ChangeAllowed(ctx); err != nil {
		return result, errors.Trace(err)
	}

	if len(args.Entities) == 0 {
		return result, nil
	}

	isSuperUser, err := api.hasControllerAdminAccess()
	if err != nil && !errors.Is(err, authentication.ErrorEntityMissingPermission) {
		return result, errors.Trace(err)
	}

	result.Results = make([]params.AddUserResult, len(args.Entities))
	for i, arg := range args.Entities {
		result.Results[i] = params.AddUserResult{Tag: arg.Tag}

		// TODO(anvial): Legacy block to delete when user domain wire up is complete.
		user, err := api.getUser(arg.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		if isSuperUser && api.apiUser != user.Tag() {
			key, err := user.ResetPassword()
			if err != nil {
				result.Results[i].Error = apiservererrors.ServerError(err)
				continue
			}
			result.Results[i].SecretKey = key
		} else {
			result.Results[i].Error = apiservererrors.ServerError(apiservererrors.ErrPerm)
		}
		// End legacy block.

		// Get User
		usr, err := api.userService.GetUserByName(ctx, arg.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}

		// Reset password
		if isSuperUser && api.apiUser != coreuser.UserTag(usr.UUID) {
			key, err := api.userService.ResetPassword(ctx, usr.UUID)
			if err != nil {
				result.Results[i].Error = apiservererrors.ServerError(err)
				continue
			}
			result.Results[i].SecretKey = key
		} else {
			result.Results[i].Error = apiservererrors.ServerError(apiservererrors.ErrPerm)
		}
	}
	return result, nil
}
