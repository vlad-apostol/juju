// Code generated by MockGen. DO NOT EDIT.
// Source: rotate.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	watcher "github.com/juju/juju/core/watcher"
	gomock "go.uber.org/mock/gomock"
)

// Mocklogger is a mock of logger interface.
type Mocklogger struct {
	ctrl     *gomock.Controller
	recorder *MockloggerMockRecorder
}

// MockloggerMockRecorder is the mock recorder for Mocklogger.
type MockloggerMockRecorder struct {
	mock *Mocklogger
}

// NewMocklogger creates a new mock instance.
func NewMocklogger(ctrl *gomock.Controller) *Mocklogger {
	mock := &Mocklogger{ctrl: ctrl}
	mock.recorder = &MockloggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocklogger) EXPECT() *MockloggerMockRecorder {
	return m.recorder
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// MockSecretBackendManagerFacade is a mock of SecretBackendManagerFacade interface.
type MockSecretBackendManagerFacade struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendManagerFacadeMockRecorder
}

// MockSecretBackendManagerFacadeMockRecorder is the mock recorder for MockSecretBackendManagerFacade.
type MockSecretBackendManagerFacadeMockRecorder struct {
	mock *MockSecretBackendManagerFacade
}

// NewMockSecretBackendManagerFacade creates a new mock instance.
func NewMockSecretBackendManagerFacade(ctrl *gomock.Controller) *MockSecretBackendManagerFacade {
	mock := &MockSecretBackendManagerFacade{ctrl: ctrl}
	mock.recorder = &MockSecretBackendManagerFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendManagerFacade) EXPECT() *MockSecretBackendManagerFacadeMockRecorder {
	return m.recorder
}

// RotateBackendTokens mocks base method.
func (m *MockSecretBackendManagerFacade) RotateBackendTokens(info ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range info {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RotateBackendTokens", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RotateBackendTokens indicates an expected call of RotateBackendTokens.
func (mr *MockSecretBackendManagerFacadeMockRecorder) RotateBackendTokens(info ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateBackendTokens", reflect.TypeOf((*MockSecretBackendManagerFacade)(nil).RotateBackendTokens), info...)
}

// WatchTokenRotationChanges mocks base method.
func (m *MockSecretBackendManagerFacade) WatchTokenRotationChanges() (watcher.SecretBackendRotateWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchTokenRotationChanges")
	ret0, _ := ret[0].(watcher.SecretBackendRotateWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchTokenRotationChanges indicates an expected call of WatchTokenRotationChanges.
func (mr *MockSecretBackendManagerFacadeMockRecorder) WatchTokenRotationChanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchTokenRotationChanges", reflect.TypeOf((*MockSecretBackendManagerFacade)(nil).WatchTokenRotationChanges))
}
