// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/sshserver (interfaces: SystemState)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/system_state_mock.go github.com/juju/juju/worker/sshserver SystemState
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	state "github.com/juju/juju/state"
	gomock "go.uber.org/mock/gomock"
)

// MockSystemState is a mock of SystemState interface.
type MockSystemState struct {
	ctrl     *gomock.Controller
	recorder *MockSystemStateMockRecorder
}

// MockSystemStateMockRecorder is the mock recorder for MockSystemState.
type MockSystemStateMockRecorder struct {
	mock *MockSystemState
}

// NewMockSystemState creates a new mock instance.
func NewMockSystemState(ctrl *gomock.Controller) *MockSystemState {
	mock := &MockSystemState{ctrl: ctrl}
	mock.recorder = &MockSystemStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemState) EXPECT() *MockSystemStateMockRecorder {
	return m.recorder
}

// WatchControllerConfig mocks base method.
func (m *MockSystemState) WatchControllerConfig() state.NotifyWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchControllerConfig")
	ret0, _ := ret[0].(state.NotifyWatcher)
	return ret0
}

// WatchControllerConfig indicates an expected call of WatchControllerConfig.
func (mr *MockSystemStateMockRecorder) WatchControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchControllerConfig", reflect.TypeOf((*MockSystemState)(nil).WatchControllerConfig))
}
