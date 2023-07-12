// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/externalcontroller/service (interfaces: State)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	crossmodel "github.com/juju/juju/core/crossmodel"
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

// Controller mocks base method.
func (m *MockState) Controller(arg0 context.Context, arg1 string) (*crossmodel.ControllerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Controller", arg0, arg1)
	ret0, _ := ret[0].(*crossmodel.ControllerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Controller indicates an expected call of Controller.
func (mr *MockStateMockRecorder) Controller(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Controller", reflect.TypeOf((*MockState)(nil).Controller), arg0, arg1)
}

// ControllerForModel mocks base method.
func (m *MockState) ControllerForModel(arg0 context.Context, arg1 string) (*crossmodel.ControllerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerForModel", arg0, arg1)
	ret0, _ := ret[0].(*crossmodel.ControllerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerForModel indicates an expected call of ControllerForModel.
func (mr *MockStateMockRecorder) ControllerForModel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerForModel", reflect.TypeOf((*MockState)(nil).ControllerForModel), arg0, arg1)
}

// UpdateExternalController mocks base method.
func (m *MockState) UpdateExternalController(arg0 context.Context, arg1 crossmodel.ControllerInfo, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExternalController", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalController indicates an expected call of UpdateExternalController.
func (mr *MockStateMockRecorder) UpdateExternalController(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalController", reflect.TypeOf((*MockState)(nil).UpdateExternalController), arg0, arg1, arg2)
}
