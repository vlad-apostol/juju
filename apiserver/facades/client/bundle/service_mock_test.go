// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/bundle (interfaces: NetworkService,ApplicationService)
//
// Generated by this command:
//
//	mockgen -typed -package bundle_test -destination service_mock_test.go github.com/juju/juju/apiserver/facades/client/bundle NetworkService,ApplicationService
//

// Package bundle_test is a generated GoMock package.
package bundle_test

import (
	context "context"
	reflect "reflect"

	charm "github.com/juju/juju/core/charm"
	network "github.com/juju/juju/core/network"
	charm0 "github.com/juju/juju/domain/application/charm"
	charm1 "github.com/juju/juju/internal/charm"
	gomock "go.uber.org/mock/gomock"
)

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// GetAllSpaces mocks base method.
func (m *MockNetworkService) GetAllSpaces(arg0 context.Context) (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpaces", arg0)
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpaces indicates an expected call of GetAllSpaces.
func (mr *MockNetworkServiceMockRecorder) GetAllSpaces(arg0 any) *MockNetworkServiceGetAllSpacesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpaces", reflect.TypeOf((*MockNetworkService)(nil).GetAllSpaces), arg0)
	return &MockNetworkServiceGetAllSpacesCall{Call: call}
}

// MockNetworkServiceGetAllSpacesCall wrap *gomock.Call
type MockNetworkServiceGetAllSpacesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockNetworkServiceGetAllSpacesCall) Return(arg0 network.SpaceInfos, arg1 error) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockNetworkServiceGetAllSpacesCall) Do(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockNetworkServiceGetAllSpacesCall) DoAndReturn(f func(context.Context) (network.SpaceInfos, error)) *MockNetworkServiceGetAllSpacesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// GetCharm mocks base method.
func (m *MockApplicationService) GetCharm(arg0 context.Context, arg1 charm.ID) (charm1.Charm, charm0.CharmLocator, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharm", arg0, arg1)
	ret0, _ := ret[0].(charm1.Charm)
	ret1, _ := ret[1].(charm0.CharmLocator)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetCharm indicates an expected call of GetCharm.
func (mr *MockApplicationServiceMockRecorder) GetCharm(arg0, arg1 any) *MockApplicationServiceGetCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharm", reflect.TypeOf((*MockApplicationService)(nil).GetCharm), arg0, arg1)
	return &MockApplicationServiceGetCharmCall{Call: call}
}

// MockApplicationServiceGetCharmCall wrap *gomock.Call
type MockApplicationServiceGetCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetCharmCall) Return(arg0 charm1.Charm, arg1 charm0.CharmLocator, arg2 bool, arg3 error) *MockApplicationServiceGetCharmCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetCharmCall) Do(f func(context.Context, charm.ID) (charm1.Charm, charm0.CharmLocator, bool, error)) *MockApplicationServiceGetCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetCharmCall) DoAndReturn(f func(context.Context, charm.ID) (charm1.Charm, charm0.CharmLocator, bool, error)) *MockApplicationServiceGetCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCharmID mocks base method.
func (m *MockApplicationService) GetCharmID(arg0 context.Context, arg1 charm0.GetCharmArgs) (charm.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmID", arg0, arg1)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmID indicates an expected call of GetCharmID.
func (mr *MockApplicationServiceMockRecorder) GetCharmID(arg0, arg1 any) *MockApplicationServiceGetCharmIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmID", reflect.TypeOf((*MockApplicationService)(nil).GetCharmID), arg0, arg1)
	return &MockApplicationServiceGetCharmIDCall{Call: call}
}

// MockApplicationServiceGetCharmIDCall wrap *gomock.Call
type MockApplicationServiceGetCharmIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetCharmIDCall) Return(arg0 charm.ID, arg1 error) *MockApplicationServiceGetCharmIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetCharmIDCall) Do(f func(context.Context, charm0.GetCharmArgs) (charm.ID, error)) *MockApplicationServiceGetCharmIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetCharmIDCall) DoAndReturn(f func(context.Context, charm0.GetCharmArgs) (charm.ID, error)) *MockApplicationServiceGetCharmIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
