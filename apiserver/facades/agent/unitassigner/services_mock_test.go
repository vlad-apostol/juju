// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/unitassigner (interfaces: ApplicationService)
//
// Generated by this command:
//
//	mockgen -typed -package unitassigner -destination services_mock_test.go github.com/juju/juju/apiserver/facades/agent/unitassigner ApplicationService
//

// Package unitassigner is a generated GoMock package.
package unitassigner

import (
	context "context"
	reflect "reflect"

	status "github.com/juju/juju/core/status"
	unit "github.com/juju/juju/core/unit"
	gomock "go.uber.org/mock/gomock"
)

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

// SetUnitAgentStatus mocks base method.
func (m *MockApplicationService) SetUnitAgentStatus(arg0 context.Context, arg1 unit.Name, arg2 *status.StatusInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnitAgentStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitAgentStatus indicates an expected call of SetUnitAgentStatus.
func (mr *MockApplicationServiceMockRecorder) SetUnitAgentStatus(arg0, arg1, arg2 any) *MockApplicationServiceSetUnitAgentStatusCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitAgentStatus", reflect.TypeOf((*MockApplicationService)(nil).SetUnitAgentStatus), arg0, arg1, arg2)
	return &MockApplicationServiceSetUnitAgentStatusCall{Call: call}
}

// MockApplicationServiceSetUnitAgentStatusCall wrap *gomock.Call
type MockApplicationServiceSetUnitAgentStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceSetUnitAgentStatusCall) Return(arg0 error) *MockApplicationServiceSetUnitAgentStatusCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceSetUnitAgentStatusCall) Do(f func(context.Context, unit.Name, *status.StatusInfo) error) *MockApplicationServiceSetUnitAgentStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceSetUnitAgentStatusCall) DoAndReturn(f func(context.Context, unit.Name, *status.StatusInfo) error) *MockApplicationServiceSetUnitAgentStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
