// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/unitstate/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination state_mock_test.go github.com/juju/juju/domain/unitstate/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	unit "github.com/juju/juju/core/unit"
	domain "github.com/juju/juju/domain"
	unitstate "github.com/juju/juju/domain/unitstate"
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

// EnsureUnitStateRecord mocks base method.
func (m *MockState) EnsureUnitStateRecord(arg0 domain.AtomicContext, arg1 unit.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureUnitStateRecord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureUnitStateRecord indicates an expected call of EnsureUnitStateRecord.
func (mr *MockStateMockRecorder) EnsureUnitStateRecord(arg0, arg1 any) *MockStateEnsureUnitStateRecordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureUnitStateRecord", reflect.TypeOf((*MockState)(nil).EnsureUnitStateRecord), arg0, arg1)
	return &MockStateEnsureUnitStateRecordCall{Call: call}
}

// MockStateEnsureUnitStateRecordCall wrap *gomock.Call
type MockStateEnsureUnitStateRecordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateEnsureUnitStateRecordCall) Return(arg0 error) *MockStateEnsureUnitStateRecordCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateEnsureUnitStateRecordCall) Do(f func(domain.AtomicContext, unit.UUID) error) *MockStateEnsureUnitStateRecordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateEnsureUnitStateRecordCall) DoAndReturn(f func(domain.AtomicContext, unit.UUID) error) *MockStateEnsureUnitStateRecordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitState mocks base method.
func (m *MockState) GetUnitState(arg0 context.Context, arg1 unit.Name) (unitstate.RetrievedUnitState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitState", arg0, arg1)
	ret0, _ := ret[0].(unitstate.RetrievedUnitState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitState indicates an expected call of GetUnitState.
func (mr *MockStateMockRecorder) GetUnitState(arg0, arg1 any) *MockStateGetUnitStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitState", reflect.TypeOf((*MockState)(nil).GetUnitState), arg0, arg1)
	return &MockStateGetUnitStateCall{Call: call}
}

// MockStateGetUnitStateCall wrap *gomock.Call
type MockStateGetUnitStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetUnitStateCall) Return(arg0 unitstate.RetrievedUnitState, arg1 error) *MockStateGetUnitStateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetUnitStateCall) Do(f func(context.Context, unit.Name) (unitstate.RetrievedUnitState, error)) *MockStateGetUnitStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetUnitStateCall) DoAndReturn(f func(context.Context, unit.Name) (unitstate.RetrievedUnitState, error)) *MockStateGetUnitStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUnitUUIDForName mocks base method.
func (m *MockState) GetUnitUUIDForName(arg0 domain.AtomicContext, arg1 unit.Name) (unit.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnitUUIDForName", arg0, arg1)
	ret0, _ := ret[0].(unit.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnitUUIDForName indicates an expected call of GetUnitUUIDForName.
func (mr *MockStateMockRecorder) GetUnitUUIDForName(arg0, arg1 any) *MockStateGetUnitUUIDForNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnitUUIDForName", reflect.TypeOf((*MockState)(nil).GetUnitUUIDForName), arg0, arg1)
	return &MockStateGetUnitUUIDForNameCall{Call: call}
}

// MockStateGetUnitUUIDForNameCall wrap *gomock.Call
type MockStateGetUnitUUIDForNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetUnitUUIDForNameCall) Return(arg0 unit.UUID, arg1 error) *MockStateGetUnitUUIDForNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetUnitUUIDForNameCall) Do(f func(domain.AtomicContext, unit.Name) (unit.UUID, error)) *MockStateGetUnitUUIDForNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetUnitUUIDForNameCall) DoAndReturn(f func(domain.AtomicContext, unit.Name) (unit.UUID, error)) *MockStateGetUnitUUIDForNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RunAtomic mocks base method.
func (m *MockState) RunAtomic(arg0 context.Context, arg1 func(domain.AtomicContext) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunAtomic", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunAtomic indicates an expected call of RunAtomic.
func (mr *MockStateMockRecorder) RunAtomic(arg0, arg1 any) *MockStateRunAtomicCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAtomic", reflect.TypeOf((*MockState)(nil).RunAtomic), arg0, arg1)
	return &MockStateRunAtomicCall{Call: call}
}

// MockStateRunAtomicCall wrap *gomock.Call
type MockStateRunAtomicCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateRunAtomicCall) Return(arg0 error) *MockStateRunAtomicCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateRunAtomicCall) Do(f func(context.Context, func(domain.AtomicContext) error) error) *MockStateRunAtomicCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateRunAtomicCall) DoAndReturn(f func(context.Context, func(domain.AtomicContext) error) error) *MockStateRunAtomicCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetUnitStateCharm mocks base method.
func (m *MockState) SetUnitStateCharm(arg0 domain.AtomicContext, arg1 unit.UUID, arg2 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnitStateCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitStateCharm indicates an expected call of SetUnitStateCharm.
func (mr *MockStateMockRecorder) SetUnitStateCharm(arg0, arg1, arg2 any) *MockStateSetUnitStateCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitStateCharm", reflect.TypeOf((*MockState)(nil).SetUnitStateCharm), arg0, arg1, arg2)
	return &MockStateSetUnitStateCharmCall{Call: call}
}

// MockStateSetUnitStateCharmCall wrap *gomock.Call
type MockStateSetUnitStateCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetUnitStateCharmCall) Return(arg0 error) *MockStateSetUnitStateCharmCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetUnitStateCharmCall) Do(f func(domain.AtomicContext, unit.UUID, map[string]string) error) *MockStateSetUnitStateCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetUnitStateCharmCall) DoAndReturn(f func(domain.AtomicContext, unit.UUID, map[string]string) error) *MockStateSetUnitStateCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetUnitStateRelation mocks base method.
func (m *MockState) SetUnitStateRelation(arg0 domain.AtomicContext, arg1 unit.UUID, arg2 map[int]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnitStateRelation", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitStateRelation indicates an expected call of SetUnitStateRelation.
func (mr *MockStateMockRecorder) SetUnitStateRelation(arg0, arg1, arg2 any) *MockStateSetUnitStateRelationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitStateRelation", reflect.TypeOf((*MockState)(nil).SetUnitStateRelation), arg0, arg1, arg2)
	return &MockStateSetUnitStateRelationCall{Call: call}
}

// MockStateSetUnitStateRelationCall wrap *gomock.Call
type MockStateSetUnitStateRelationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetUnitStateRelationCall) Return(arg0 error) *MockStateSetUnitStateRelationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetUnitStateRelationCall) Do(f func(domain.AtomicContext, unit.UUID, map[int]string) error) *MockStateSetUnitStateRelationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetUnitStateRelationCall) DoAndReturn(f func(domain.AtomicContext, unit.UUID, map[int]string) error) *MockStateSetUnitStateRelationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUnitStateSecret mocks base method.
func (m *MockState) UpdateUnitStateSecret(arg0 domain.AtomicContext, arg1 unit.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnitStateSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUnitStateSecret indicates an expected call of UpdateUnitStateSecret.
func (mr *MockStateMockRecorder) UpdateUnitStateSecret(arg0, arg1, arg2 any) *MockStateUpdateUnitStateSecretCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnitStateSecret", reflect.TypeOf((*MockState)(nil).UpdateUnitStateSecret), arg0, arg1, arg2)
	return &MockStateUpdateUnitStateSecretCall{Call: call}
}

// MockStateUpdateUnitStateSecretCall wrap *gomock.Call
type MockStateUpdateUnitStateSecretCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateUpdateUnitStateSecretCall) Return(arg0 error) *MockStateUpdateUnitStateSecretCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateUpdateUnitStateSecretCall) Do(f func(domain.AtomicContext, unit.UUID, string) error) *MockStateUpdateUnitStateSecretCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateUpdateUnitStateSecretCall) DoAndReturn(f func(domain.AtomicContext, unit.UUID, string) error) *MockStateUpdateUnitStateSecretCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUnitStateStorage mocks base method.
func (m *MockState) UpdateUnitStateStorage(arg0 domain.AtomicContext, arg1 unit.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnitStateStorage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUnitStateStorage indicates an expected call of UpdateUnitStateStorage.
func (mr *MockStateMockRecorder) UpdateUnitStateStorage(arg0, arg1, arg2 any) *MockStateUpdateUnitStateStorageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnitStateStorage", reflect.TypeOf((*MockState)(nil).UpdateUnitStateStorage), arg0, arg1, arg2)
	return &MockStateUpdateUnitStateStorageCall{Call: call}
}

// MockStateUpdateUnitStateStorageCall wrap *gomock.Call
type MockStateUpdateUnitStateStorageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateUpdateUnitStateStorageCall) Return(arg0 error) *MockStateUpdateUnitStateStorageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateUpdateUnitStateStorageCall) Do(f func(domain.AtomicContext, unit.UUID, string) error) *MockStateUpdateUnitStateStorageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateUpdateUnitStateStorageCall) DoAndReturn(f func(domain.AtomicContext, unit.UUID, string) error) *MockStateUpdateUnitStateStorageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUnitStateUniter mocks base method.
func (m *MockState) UpdateUnitStateUniter(arg0 domain.AtomicContext, arg1 unit.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnitStateUniter", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUnitStateUniter indicates an expected call of UpdateUnitStateUniter.
func (mr *MockStateMockRecorder) UpdateUnitStateUniter(arg0, arg1, arg2 any) *MockStateUpdateUnitStateUniterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnitStateUniter", reflect.TypeOf((*MockState)(nil).UpdateUnitStateUniter), arg0, arg1, arg2)
	return &MockStateUpdateUnitStateUniterCall{Call: call}
}

// MockStateUpdateUnitStateUniterCall wrap *gomock.Call
type MockStateUpdateUnitStateUniterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateUpdateUnitStateUniterCall) Return(arg0 error) *MockStateUpdateUnitStateUniterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateUpdateUnitStateUniterCall) Do(f func(domain.AtomicContext, unit.UUID, string) error) *MockStateUpdateUnitStateUniterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateUpdateUnitStateUniterCall) DoAndReturn(f func(domain.AtomicContext, unit.UUID, string) error) *MockStateUpdateUnitStateUniterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
