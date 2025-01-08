// Copyright 2019 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package instancemutater_test

import (
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v4"
	"github.com/juju/worker/v4/workertest"
	"go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/apiserver/facades/agent/instancemutater"
	"github.com/juju/juju/apiserver/facades/agent/instancemutater/mocks"
	"github.com/juju/juju/core/charm"
	"github.com/juju/juju/core/machine"
	"github.com/juju/juju/core/watcher/watchertest"
	applicationcharm "github.com/juju/juju/domain/application/charm"
	applicationerrors "github.com/juju/juju/domain/application/errors"
	internalcharm "github.com/juju/juju/internal/charm"
	"github.com/juju/juju/state"
)

type lxdProfileWatcherSuite struct {
	state     *mocks.MockInstanceMutaterState
	machine0  *mocks.MockMachine
	unit      *mocks.MockUnit
	principal *mocks.MockUnit
	app       *mocks.MockApplication

	charmsWatcher      *mocks.MockStringsWatcher
	appWatcher         *mocks.MockStringsWatcher
	unitsWatcher       *mocks.MockStringsWatcher
	instanceWatcher    *mocks.MockNotifyWatcher
	machineService     *mocks.MockMachineService
	applicationService *mocks.MockApplicationService

	charmChanges    chan []string
	appChanges      chan []string
	unitChanges     chan []string
	instanceChanges chan struct{}

	wc0 watchertest.NotifyWatcherC
}

var _ = gc.Suite(&lxdProfileWatcherSuite{})

func (s *lxdProfileWatcherSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.state = mocks.NewMockInstanceMutaterState(ctrl)
	s.unit = mocks.NewMockUnit(ctrl)
	s.principal = mocks.NewMockUnit(ctrl)
	s.app = mocks.NewMockApplication(ctrl)
	s.machine0 = mocks.NewMockMachine(ctrl)

	s.charmsWatcher = mocks.NewMockStringsWatcher(ctrl)
	s.appWatcher = mocks.NewMockStringsWatcher(ctrl)
	s.unitsWatcher = mocks.NewMockStringsWatcher(ctrl)
	s.instanceWatcher = mocks.NewMockNotifyWatcher(ctrl)
	s.machineService = mocks.NewMockMachineService(ctrl)
	s.applicationService = mocks.NewMockApplicationService(ctrl)

	s.charmChanges = make(chan []string, 1)
	s.appChanges = make(chan []string, 1)
	s.unitChanges = make(chan []string, 1)
	s.instanceChanges = make(chan struct{}, 1)

	return ctrl
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherStartStop(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherNoProfile(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioNoProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.setupPrincipalUnit()
	s.unitChanges <- []string{"foo/0"}
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherProfile(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioNoExistingUnitsWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.setupPrincipalUnit()
	curlStr := "ch:name-me"
	s.unit.EXPECT().CharmURL().Return(&curlStr)
	s.unitChanges <- []string{"foo/0"}
	s.wc0.AssertOneChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherNewCharmRev(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	// Start with a charm having a profile so change the charm's profile
	// from existing to not, should be notified to remove the profile.
	s.updateCharmForMachineLXDProfileWatcherWithoutProfile(c, "2")
	s.wc0.AssertOneChange()

	// Changing the charm url, and the profile stays empty,
	// should not be notified to remove the profile.
	s.updateCharmForMachineLXDProfileWatcherWithoutProfile(c, "3")
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherCharmMetadataChange(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	// Start with a charm not having a profile.
	s.updateCharmForMachineLXDProfileWatcherWithoutProfile(c, "2")
	s.wc0.AssertOneChange()

	// Simulate an asynchronous charm download scenario where the downloaded
	// charm specifies an LXD profile. This should trigger a change.
	s.updateCharmForMachineLXDProfileWatcherWithProfile(c, "2")
	s.wc0.AssertOneChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherAddUnit(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioNoExistingUnitsWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	// New unit added to existing machine doesn't have a charm url yet.
	// It may have been added without a machine id either.

	s.state.EXPECT().Unit("bar/0").Return(s.unit, nil)
	s.unit.EXPECT().Life().Return(state.Alive)
	s.unit.EXPECT().PrincipalName().Return("", false)
	s.unit.EXPECT().AssignedMachineId().Return("", errors.NotAssignedf(""))
	s.unitChanges <- []string{"bar/0"}
	s.wc0.AssertNoChange()

	// Add the machine id, this time we should get a notification.
	s.state.EXPECT().Unit("bar/0").Return(s.unit, nil)
	s.unit.EXPECT().Life().Return(state.Alive)
	s.unit.EXPECT().PrincipalName().Return("", false)
	s.unit.EXPECT().AssignedMachineId().Return("0", nil)
	curlStr := "ch:name-me"
	s.unit.EXPECT().CharmURL().Return(&curlStr)
	s.unitChanges <- []string{"bar/0"}
	s.wc0.AssertOneChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherAddUnitWrongMachine(c *gc.C) {
	defer s.setup(c).Finish()

	s.unit.EXPECT().ApplicationName().AnyTimes().Return("foo")
	s.unit.EXPECT().Name().AnyTimes().Return("foo/0")
	s.machine0.EXPECT().Units().Return(nil, nil)

	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.state.EXPECT().Unit("foo/0").Return(s.unit, nil)
	s.unit.EXPECT().Life().Return(state.Alive)
	s.unit.EXPECT().AssignedMachineId().Return("1", nil)
	s.unit.EXPECT().PrincipalName().Return("", false)
	s.unitChanges <- []string{"foo/0"}
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherSubordinateWithProfile(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioNoExistingUnitsWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.assertAddSubordinate()
	s.wc0.AssertOneChange()
}

func (s *lxdProfileWatcherSuite) assertAddSubordinate() {
	// Add a new subordinate unit with a profile of a new application.

	s.state.EXPECT().Unit("foo/0").Return(s.unit, nil)
	s.unit.EXPECT().Life().Return(state.Alive)
	s.unit.EXPECT().PrincipalName().Return("principal/0", true)
	s.unit.EXPECT().AssignedMachineId().Return("0", nil)

	curlStr := "ch:name-me"
	s.unit.EXPECT().CharmURL().Return(&curlStr)
	s.unitChanges <- []string{"foo/0"}
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherSubordinateWithProfileUpdateUnit(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.setupScenarioNoExistingUnitsWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.assertAddSubordinate()
	s.wc0.AssertOneChange()

	// Add another subordinate and expect a change.
	another := mocks.NewMockUnit(ctrl)
	another.EXPECT().Name().AnyTimes().Return("foo/1")
	another.EXPECT().ApplicationName().AnyTimes().Return("foo")
	another.EXPECT().Life().Return(state.Alive)
	another.EXPECT().PrincipalName().Return("principal/0", true)
	another.EXPECT().AssignedMachineId().Return("0", nil)
	s.state.EXPECT().Unit("foo/1").Return(another, nil)

	s.unitChanges <- []string{"foo/1"}
	s.wc0.AssertOneChange()

	// A general change for an existing unit should cause no notification.
	another.EXPECT().Life().Return(state.Alive)
	another.EXPECT().PrincipalName().Return("principal/0", true)
	another.EXPECT().AssignedMachineId().Return("0", nil)
	s.state.EXPECT().Unit("foo/1").Return(another, nil)
	s.unitChanges <- []string{"foo/1"}
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherSubordinateNoProfile(c *gc.C) {
	defer s.setup(c).Finish()

	s.unit.EXPECT().ApplicationName().AnyTimes().Return("foo")
	s.unit.EXPECT().Name().AnyTimes().Return("foo/0")

	curl := "ch:name-me"
	s.machine0.EXPECT().Units().Return(nil, nil)
	s.assertCharmWithoutLXDProfile(c, curl)

	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.assertAddSubordinate()
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherRemoveUnitWithProfileTwoUnits(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.setupScenarioWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.state.EXPECT().Unit("foo/0").Return(s.unit, nil)
	s.wc0.AssertNoChange()

	// Add a new unit.
	another := mocks.NewMockUnit(ctrl)
	another.EXPECT().Name().AnyTimes().Return("foo/1")
	another.EXPECT().ApplicationName().AnyTimes().Return("foo")
	another.EXPECT().Life().Return(state.Alive)
	another.EXPECT().PrincipalName().Return("", false)
	another.EXPECT().AssignedMachineId().Return("0", nil)
	s.state.EXPECT().Unit("foo/1").Return(another, nil)

	s.unitChanges <- []string{"foo/1"}
	s.wc0.AssertOneChange()

	// Remove the original unit.
	s.unit.EXPECT().Life().Return(state.Dead)
	s.unitChanges <- []string{"foo/0"}
	s.wc0.AssertOneChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherRemoveOnlyUnit(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.setupScenarioNoExistingUnitsWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.wc0.AssertNoChange()

	s.setupPrincipalUnit()
	curlStr := "ch:name-me"
	s.unit.EXPECT().CharmURL().Return(&curlStr)
	s.unitChanges <- []string{"foo/0"}
	s.wc0.AssertOneChange()

	// Remove the original unit.
	s.state.EXPECT().Unit("foo/0").Return(s.unit, nil)
	s.unit.EXPECT().Life().Return(state.Dead)
	s.unitChanges <- []string{"foo/0"}
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherRemoveUnitWrongMachine(c *gc.C) {
	ctrl := s.setup(c)
	defer ctrl.Finish()

	s.unit.EXPECT().ApplicationName().AnyTimes().Return("foo")
	s.unit.EXPECT().Name().AnyTimes().Return("foo/0")
	s.machine0.EXPECT().Units().Return(nil, nil)

	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.wc0.AssertNoChange()

	s.state.EXPECT().Unit("bar/0").Return(s.unit, nil)
	s.unit.EXPECT().Life().Return(state.Dead)
	s.unitChanges <- []string{"bar/0"}
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherAppChangeCharmURLNotFound(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioWithProfile(c)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.updateCharmForMachineLXDProfileWatcherWithoutProfile(c, "2")
	s.wc0.AssertOneChange()

	s.state.EXPECT().Application("foo").Return(s.app, nil)
	curl := "ch:name-me-3"
	s.app.EXPECT().CharmURL().Return(&curl)
	s.assertCharmNotFound(c, curl)

	s.appChanges <- []string{"foo"}
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherUnitChangeAppNotFound(c *gc.C) {
	defer s.setup(c).Finish()

	s.unit.EXPECT().ApplicationName().AnyTimes().Return("foo")
	s.unit.EXPECT().Name().AnyTimes().Return("foo/0")
	s.machine0.EXPECT().Units().Return(nil, nil)

	s.setupPrincipalUnit()
	s.unit.EXPECT().CharmURL().Return(nil)
	s.unit.EXPECT().Application().Return(nil, errors.NotFoundf(""))

	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.unitChanges <- []string{"foo/0"}
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherUnitChangeCharmURLNotFound(c *gc.C) {
	defer s.setup(c).Finish()

	s.unit.EXPECT().ApplicationName().AnyTimes().Return("foo")
	s.unit.EXPECT().Name().AnyTimes().Return("foo/0")
	s.machine0.EXPECT().Units().Return(nil, nil)

	s.setupPrincipalUnit()
	curl := "ch:name-me"
	s.unit.EXPECT().CharmURL().Return(&curl)
	s.applicationService.EXPECT().GetCharmID(gomock.Any(), applicationcharm.GetCharmArgs{
		Source:   applicationcharm.CharmHubSource,
		Name:     "name-me",
		Revision: ptr(-1),
	}).Return(charm.ID(""), applicationerrors.CharmNotFound)

	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.unitChanges <- []string{"foo/0"}
	s.wc0.AssertNoChange()
}

func (s *lxdProfileWatcherSuite) TestMachineLXDProfileWatcherMachineProvisioned(c *gc.C) {
	defer s.setup(c).Finish()

	s.setupScenarioWithProfile(c)
	s.machine0.EXPECT().Id().Return("0")
	s.machineService.EXPECT().GetMachineUUID(gomock.Any(), machine.Name("0")).Return("uuid0", nil)
	s.machineService.EXPECT().InstanceID(gomock.Any(), "uuid0").Return("0", nil)
	defer workertest.CleanKill(c, s.assertStartLxdProfileWatcher(c))

	s.instanceChanges <- struct{}{}
	s.wc0.AssertOneChange()
}

func (s *lxdProfileWatcherSuite) updateCharmForMachineLXDProfileWatcherWithProfile(c *gc.C, rev string) {
	curl := "ch:name-me-" + rev
	s.state.EXPECT().Application("foo").Return(s.app, nil)
	s.app.EXPECT().CharmURL().Return(&curl)
	s.assertCharmWithLXDProfile(c, curl)
	s.charmChanges <- []string{curl}
	s.appChanges <- []string{"foo"}
}

func (s *lxdProfileWatcherSuite) updateCharmForMachineLXDProfileWatcherWithoutProfile(c *gc.C, rev string) {
	curl := "ch:name-me-" + rev
	s.state.EXPECT().Application("foo").Return(s.app, nil)
	s.app.EXPECT().CharmURL().Return(&curl)
	s.assertCharmWithoutLXDProfile(c, curl)
	s.charmChanges <- []string{curl}
	s.appChanges <- []string{"foo"}
}
func (s *lxdProfileWatcherSuite) setupWatchers(c *gc.C) {
	s.state.EXPECT().WatchApplicationCharms().Return(s.appWatcher)
	s.state.EXPECT().WatchUnits().Return(s.unitsWatcher)

	s.machineService.EXPECT().GetMachineUUID(gomock.Any(), machine.Name("0")).Return("uuid0", nil)
	s.machineService.EXPECT().WatchLXDProfiles(gomock.Any(), "uuid0").Return(s.instanceWatcher, nil)

	s.applicationService.EXPECT().WatchCharms().Return(s.charmsWatcher, nil)

	s.charmsWatcher.EXPECT().Changes().AnyTimes().Return(s.charmChanges)
	s.charmsWatcher.EXPECT().Wait().Return(nil)
	s.appWatcher.EXPECT().Changes().AnyTimes().Return(s.appChanges)
	s.appWatcher.EXPECT().Wait().Return(nil)
	s.unitsWatcher.EXPECT().Changes().AnyTimes().Return(s.unitChanges)
	s.unitsWatcher.EXPECT().Wait().Return(nil)
	s.instanceWatcher.EXPECT().Changes().AnyTimes().Return(s.instanceChanges)
	s.instanceWatcher.EXPECT().Wait().Return(nil)
}

func (s *lxdProfileWatcherSuite) assertStartLxdProfileWatcher(c *gc.C) worker.Worker {
	s.setupWatchers(c)

	s.machine0.EXPECT().Id().AnyTimes().Return("0")

	w := instancemutater.NewTestLxdProfileWatcher(c, s.machine0, s.state, s.machineService, s.applicationService)
	wc := watchertest.NewNotifyWatcherC(c, w)
	// Sends initial event.
	wc.AssertOneChange()
	s.wc0 = wc
	return w
}

func (s *lxdProfileWatcherSuite) assertCharmNotFound(c *gc.C, chURLStr string) {
	curl, err := internalcharm.ParseURL(chURLStr)
	c.Assert(err, jc.ErrorIsNil)
	source, err := applicationcharm.ParseCharmSchema(internalcharm.Schema(curl.Schema))
	c.Assert(err, jc.ErrorIsNil)

	s.applicationService.EXPECT().GetCharmID(gomock.Any(), applicationcharm.GetCharmArgs{
		Source:   source,
		Name:     "name-me",
		Revision: ptr(curl.Revision),
	}).Return(charm.ID(""), applicationerrors.CharmNotFound)
}

func (s *lxdProfileWatcherSuite) assertCharmWithLXDProfile(c *gc.C, chURLStr string) {
	curl, err := internalcharm.ParseURL(chURLStr)
	c.Assert(err, jc.ErrorIsNil)
	source, err := applicationcharm.ParseCharmSchema(internalcharm.Schema(curl.Schema))
	c.Assert(err, jc.ErrorIsNil)

	s.applicationService.EXPECT().GetCharmID(gomock.Any(), applicationcharm.GetCharmArgs{
		Source:   source,
		Name:     curl.Name,
		Revision: ptr(curl.Revision),
	}).Return(charm.ID("foo"), nil)
	s.applicationService.EXPECT().GetCharmLXDProfile(gomock.Any(), charm.ID("foo")).
		Return(internalcharm.LXDProfile{
			Config: map[string]string{"key1": "value1"},
		}, 0, nil)
}

func (s *lxdProfileWatcherSuite) assertCharmWithoutLXDProfile(c *gc.C, chURLStr string) {
	curl, err := internalcharm.ParseURL(chURLStr)
	c.Assert(err, jc.ErrorIsNil)
	source, err := applicationcharm.ParseCharmSchema(internalcharm.Schema(curl.Schema))
	c.Assert(err, jc.ErrorIsNil)

	s.applicationService.EXPECT().GetCharmID(gomock.Any(), applicationcharm.GetCharmArgs{
		Source:   source,
		Name:     curl.Name,
		Revision: ptr(curl.Revision),
	}).Return(charm.ID("foo"), nil)
	s.applicationService.EXPECT().GetCharmLXDProfile(gomock.Any(), charm.ID("foo")).
		Return(internalcharm.LXDProfile{}, 0, nil)
}

func (s *lxdProfileWatcherSuite) setupScenarioNoProfile(c *gc.C) {
	s.unit.EXPECT().ApplicationName().AnyTimes().Return("foo")
	s.unit.EXPECT().Name().AnyTimes().Return("foo/0")
	curl := "ch:name-me"
	s.machine0.EXPECT().Units().Return([]instancemutater.Unit{s.unit}, nil)
	s.unit.EXPECT().Application().Return(s.app, nil)
	s.app.EXPECT().CharmURL().Return(&curl)
	s.assertCharmWithoutLXDProfile(c, curl)
}

func (s *lxdProfileWatcherSuite) setupScenarioWithProfile(c *gc.C) {
	s.unit.EXPECT().ApplicationName().AnyTimes().Return("foo")
	s.unit.EXPECT().Name().AnyTimes().Return("foo/0")
	curl := "ch:name-me"
	s.machine0.EXPECT().Units().Return([]instancemutater.Unit{s.unit}, nil)
	s.unit.EXPECT().Application().Return(s.app, nil)
	s.app.EXPECT().CharmURL().Return(&curl)
	s.assertCharmWithLXDProfile(c, curl)
}

func (s *lxdProfileWatcherSuite) setupScenarioNoExistingUnitsWithProfile(c *gc.C) {
	s.unit.EXPECT().ApplicationName().AnyTimes().Return("foo")
	s.unit.EXPECT().Name().AnyTimes().Return("foo/0")
	s.machine0.EXPECT().Units().Return(nil, nil)
	curl := "ch:name-me"
	s.assertCharmWithLXDProfile(c, curl)
}

func (s *lxdProfileWatcherSuite) setupPrincipalUnit() {
	s.state.EXPECT().Unit("foo/0").Return(s.unit, nil)
	s.unit.EXPECT().Life().Return(state.Alive)
	s.unit.EXPECT().AssignedMachineId().Return("0", nil)
	s.unit.EXPECT().PrincipalName().Return("", false)
}
