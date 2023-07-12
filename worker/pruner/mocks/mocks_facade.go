// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/pruner (interfaces: Facade)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	watcher "github.com/juju/juju/core/watcher"
	config "github.com/juju/juju/environs/config"
	gomock "go.uber.org/mock/gomock"
)

// MockFacade is a mock of Facade interface.
type MockFacade struct {
	ctrl     *gomock.Controller
	recorder *MockFacadeMockRecorder
}

// MockFacadeMockRecorder is the mock recorder for MockFacade.
type MockFacadeMockRecorder struct {
	mock *MockFacade
}

// NewMockFacade creates a new mock instance.
func NewMockFacade(ctrl *gomock.Controller) *MockFacade {
	mock := &MockFacade{ctrl: ctrl}
	mock.recorder = &MockFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacade) EXPECT() *MockFacadeMockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockFacade) ModelConfig() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockFacadeMockRecorder) ModelConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockFacade)(nil).ModelConfig))
}

// Prune mocks base method.
func (m *MockFacade) Prune(arg0 time.Duration, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prune", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prune indicates an expected call of Prune.
func (mr *MockFacadeMockRecorder) Prune(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prune", reflect.TypeOf((*MockFacade)(nil).Prune), arg0, arg1)
}

// WatchForModelConfigChanges mocks base method.
func (m *MockFacade) WatchForModelConfigChanges() (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchForModelConfigChanges")
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchForModelConfigChanges indicates an expected call of WatchForModelConfigChanges.
func (mr *MockFacadeMockRecorder) WatchForModelConfigChanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchForModelConfigChanges", reflect.TypeOf((*MockFacade)(nil).WatchForModelConfigChanges))
}
