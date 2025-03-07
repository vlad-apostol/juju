// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/modelcmd (interfaces: SessionLoginFactory)
//
// Generated by this command:
//
//	mockgen -package user_test -destination utils_sessionloginfactory_mock_test.go github.com/juju/juju/cmd/modelcmd SessionLoginFactory
//

// Package user_test is a generated GoMock package.
package user_test

import (
	io "io"
	reflect "reflect"

	api "github.com/juju/juju/api"
	gomock "go.uber.org/mock/gomock"
)

// MockSessionLoginFactory is a mock of SessionLoginFactory interface.
type MockSessionLoginFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSessionLoginFactoryMockRecorder
}

// MockSessionLoginFactoryMockRecorder is the mock recorder for MockSessionLoginFactory.
type MockSessionLoginFactoryMockRecorder struct {
	mock *MockSessionLoginFactory
}

// NewMockSessionLoginFactory creates a new mock instance.
func NewMockSessionLoginFactory(ctrl *gomock.Controller) *MockSessionLoginFactory {
	mock := &MockSessionLoginFactory{ctrl: ctrl}
	mock.recorder = &MockSessionLoginFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionLoginFactory) EXPECT() *MockSessionLoginFactoryMockRecorder {
	return m.recorder
}

// NewLoginProvider mocks base method.
func (m *MockSessionLoginFactory) NewLoginProvider(arg0 string, arg1 io.Writer, arg2 func(string)) api.LoginProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewLoginProvider", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.LoginProvider)
	return ret0
}

// NewLoginProvider indicates an expected call of NewLoginProvider.
func (mr *MockSessionLoginFactoryMockRecorder) NewLoginProvider(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLoginProvider", reflect.TypeOf((*MockSessionLoginFactory)(nil).NewLoginProvider), arg0, arg1, arg2)
}
