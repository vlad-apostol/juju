// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/removal/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination package_mock_test.go github.com/juju/juju/domain/removal/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// RelationAdvanceLifeAndScheduleRemoval mocks base method.
func (m *MockState) RelationAdvanceLifeAndScheduleRemoval(arg0 context.Context, arg1, arg2 string, arg3 bool, arg4 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationAdvanceLifeAndScheduleRemoval", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelationAdvanceLifeAndScheduleRemoval indicates an expected call of RelationAdvanceLifeAndScheduleRemoval.
func (mr *MockStateMockRecorder) RelationAdvanceLifeAndScheduleRemoval(arg0, arg1, arg2, arg3, arg4 any) *MockStateRelationAdvanceLifeAndScheduleRemovalCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationAdvanceLifeAndScheduleRemoval", reflect.TypeOf((*MockState)(nil).RelationAdvanceLifeAndScheduleRemoval), arg0, arg1, arg2, arg3, arg4)
	return &MockStateRelationAdvanceLifeAndScheduleRemovalCall{Call: call}
}

// MockStateRelationAdvanceLifeAndScheduleRemovalCall wrap *gomock.Call
type MockStateRelationAdvanceLifeAndScheduleRemovalCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateRelationAdvanceLifeAndScheduleRemovalCall) Return(arg0 error) *MockStateRelationAdvanceLifeAndScheduleRemovalCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateRelationAdvanceLifeAndScheduleRemovalCall) Do(f func(context.Context, string, string, bool, time.Time) error) *MockStateRelationAdvanceLifeAndScheduleRemovalCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateRelationAdvanceLifeAndScheduleRemovalCall) DoAndReturn(f func(context.Context, string, string, bool, time.Time) error) *MockStateRelationAdvanceLifeAndScheduleRemovalCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RelationExists mocks base method.
func (m *MockState) RelationExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RelationExists indicates an expected call of RelationExists.
func (mr *MockStateMockRecorder) RelationExists(arg0, arg1 any) *MockStateRelationExistsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationExists", reflect.TypeOf((*MockState)(nil).RelationExists), arg0, arg1)
	return &MockStateRelationExistsCall{Call: call}
}

// MockStateRelationExistsCall wrap *gomock.Call
type MockStateRelationExistsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateRelationExistsCall) Return(arg0 bool, arg1 error) *MockStateRelationExistsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateRelationExistsCall) Do(f func(context.Context, string) (bool, error)) *MockStateRelationExistsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateRelationExistsCall) DoAndReturn(f func(context.Context, string) (bool, error)) *MockStateRelationExistsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
