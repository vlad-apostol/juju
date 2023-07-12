// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/lease/service (interfaces: State)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	lease "github.com/juju/juju/core/lease"
	utils "github.com/juju/utils/v3"
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

// ClaimLease mocks base method.
func (m *MockState) ClaimLease(arg0 context.Context, arg1 utils.UUID, arg2 lease.Key, arg3 lease.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimLease", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClaimLease indicates an expected call of ClaimLease.
func (mr *MockStateMockRecorder) ClaimLease(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimLease", reflect.TypeOf((*MockState)(nil).ClaimLease), arg0, arg1, arg2, arg3)
}

// ExpireLeases mocks base method.
func (m *MockState) ExpireLeases(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireLeases", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExpireLeases indicates an expected call of ExpireLeases.
func (mr *MockStateMockRecorder) ExpireLeases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireLeases", reflect.TypeOf((*MockState)(nil).ExpireLeases), arg0)
}

// ExtendLease mocks base method.
func (m *MockState) ExtendLease(arg0 context.Context, arg1 lease.Key, arg2 lease.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendLease", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExtendLease indicates an expected call of ExtendLease.
func (mr *MockStateMockRecorder) ExtendLease(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendLease", reflect.TypeOf((*MockState)(nil).ExtendLease), arg0, arg1, arg2)
}

// LeaseGroup mocks base method.
func (m *MockState) LeaseGroup(arg0 context.Context, arg1, arg2 string) (map[lease.Key]lease.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaseGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[lease.Key]lease.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaseGroup indicates an expected call of LeaseGroup.
func (mr *MockStateMockRecorder) LeaseGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaseGroup", reflect.TypeOf((*MockState)(nil).LeaseGroup), arg0, arg1, arg2)
}

// Leases mocks base method.
func (m *MockState) Leases(arg0 context.Context, arg1 ...lease.Key) (map[lease.Key]lease.Info, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Leases", varargs...)
	ret0, _ := ret[0].(map[lease.Key]lease.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Leases indicates an expected call of Leases.
func (mr *MockStateMockRecorder) Leases(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Leases", reflect.TypeOf((*MockState)(nil).Leases), varargs...)
}

// PinLease mocks base method.
func (m *MockState) PinLease(arg0 context.Context, arg1 lease.Key, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PinLease", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PinLease indicates an expected call of PinLease.
func (mr *MockStateMockRecorder) PinLease(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PinLease", reflect.TypeOf((*MockState)(nil).PinLease), arg0, arg1, arg2)
}

// Pinned mocks base method.
func (m *MockState) Pinned(arg0 context.Context) (map[lease.Key][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pinned", arg0)
	ret0, _ := ret[0].(map[lease.Key][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pinned indicates an expected call of Pinned.
func (mr *MockStateMockRecorder) Pinned(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pinned", reflect.TypeOf((*MockState)(nil).Pinned), arg0)
}

// RevokeLease mocks base method.
func (m *MockState) RevokeLease(arg0 context.Context, arg1 lease.Key, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeLease", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeLease indicates an expected call of RevokeLease.
func (mr *MockStateMockRecorder) RevokeLease(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeLease", reflect.TypeOf((*MockState)(nil).RevokeLease), arg0, arg1, arg2)
}

// UnpinLease mocks base method.
func (m *MockState) UnpinLease(arg0 context.Context, arg1 lease.Key, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpinLease", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpinLease indicates an expected call of UnpinLease.
func (mr *MockStateMockRecorder) UnpinLease(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpinLease", reflect.TypeOf((*MockState)(nil).UnpinLease), arg0, arg1, arg2)
}
