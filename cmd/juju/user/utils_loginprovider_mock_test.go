// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/api (interfaces: LoginProvider)
//
// Generated by this command:
//
//	mockgen -package user_test -destination utils_loginprovider_mock_test.go github.com/juju/juju/api LoginProvider
//

// Package user_test is a generated GoMock package.
package user_test

import (
	context "context"
	http "net/http"
	reflect "reflect"

	api "github.com/juju/juju/api"
	base "github.com/juju/juju/api/base"
	gomock "go.uber.org/mock/gomock"
)

// MockLoginProvider is a mock of LoginProvider interface.
type MockLoginProvider struct {
	ctrl     *gomock.Controller
	recorder *MockLoginProviderMockRecorder
}

// MockLoginProviderMockRecorder is the mock recorder for MockLoginProvider.
type MockLoginProviderMockRecorder struct {
	mock *MockLoginProvider
}

// NewMockLoginProvider creates a new mock instance.
func NewMockLoginProvider(ctrl *gomock.Controller) *MockLoginProvider {
	mock := &MockLoginProvider{ctrl: ctrl}
	mock.recorder = &MockLoginProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginProvider) EXPECT() *MockLoginProviderMockRecorder {
	return m.recorder
}

// AuthHeader mocks base method.
func (m *MockLoginProvider) AuthHeader() (http.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthHeader")
	ret0, _ := ret[0].(http.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthHeader indicates an expected call of AuthHeader.
func (mr *MockLoginProviderMockRecorder) AuthHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthHeader", reflect.TypeOf((*MockLoginProvider)(nil).AuthHeader))
}

// Login mocks base method.
func (m *MockLoginProvider) Login(arg0 context.Context, arg1 base.APICaller) (*api.LoginResultParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*api.LoginResultParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockLoginProviderMockRecorder) Login(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockLoginProvider)(nil).Login), arg0, arg1)
}
