// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/provisioner (interfaces: Machine,BridgePolicy,Unit,Application,Charm)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/package_mock.go github.com/juju/juju/apiserver/facades/agent/provisioner Machine,BridgePolicy,Unit,Application,Charm
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	charm "github.com/juju/juju/internal/charm"
	set "github.com/juju/collections/set"
	provisioner "github.com/juju/juju/apiserver/facades/agent/provisioner"
	constraints "github.com/juju/juju/core/constraints"
	instance "github.com/juju/juju/core/instance"
	network "github.com/juju/juju/core/network"
	network0 "github.com/juju/juju/internal/network"
	containerizer "github.com/juju/juju/internal/network/containerizer"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// AllDeviceAddresses mocks base method.
func (m *MockMachine) AllDeviceAddresses() ([]containerizer.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllDeviceAddresses")
	ret0, _ := ret[0].([]containerizer.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllDeviceAddresses indicates an expected call of AllDeviceAddresses.
func (mr *MockMachineMockRecorder) AllDeviceAddresses() *MockMachineAllDeviceAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDeviceAddresses", reflect.TypeOf((*MockMachine)(nil).AllDeviceAddresses))
	return &MockMachineAllDeviceAddressesCall{Call: call}
}

// MockMachineAllDeviceAddressesCall wrap *gomock.Call
type MockMachineAllDeviceAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineAllDeviceAddressesCall) Return(arg0 []containerizer.Address, arg1 error) *MockMachineAllDeviceAddressesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineAllDeviceAddressesCall) Do(f func() ([]containerizer.Address, error)) *MockMachineAllDeviceAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineAllDeviceAddressesCall) DoAndReturn(f func() ([]containerizer.Address, error)) *MockMachineAllDeviceAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllLinkLayerDevices mocks base method.
func (m *MockMachine) AllLinkLayerDevices() ([]containerizer.LinkLayerDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllLinkLayerDevices")
	ret0, _ := ret[0].([]containerizer.LinkLayerDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllLinkLayerDevices indicates an expected call of AllLinkLayerDevices.
func (mr *MockMachineMockRecorder) AllLinkLayerDevices() *MockMachineAllLinkLayerDevicesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllLinkLayerDevices", reflect.TypeOf((*MockMachine)(nil).AllLinkLayerDevices))
	return &MockMachineAllLinkLayerDevicesCall{Call: call}
}

// MockMachineAllLinkLayerDevicesCall wrap *gomock.Call
type MockMachineAllLinkLayerDevicesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineAllLinkLayerDevicesCall) Return(arg0 []containerizer.LinkLayerDevice, arg1 error) *MockMachineAllLinkLayerDevicesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineAllLinkLayerDevicesCall) Do(f func() ([]containerizer.LinkLayerDevice, error)) *MockMachineAllLinkLayerDevicesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineAllLinkLayerDevicesCall) DoAndReturn(f func() ([]containerizer.LinkLayerDevice, error)) *MockMachineAllLinkLayerDevicesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AllSpaces mocks base method.
func (m *MockMachine) AllSpaces(arg0 network.SubnetInfos) (set.Strings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllSpaces", arg0)
	ret0, _ := ret[0].(set.Strings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllSpaces indicates an expected call of AllSpaces.
func (mr *MockMachineMockRecorder) AllSpaces(arg0 any) *MockMachineAllSpacesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSpaces", reflect.TypeOf((*MockMachine)(nil).AllSpaces), arg0)
	return &MockMachineAllSpacesCall{Call: call}
}

// MockMachineAllSpacesCall wrap *gomock.Call
type MockMachineAllSpacesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineAllSpacesCall) Return(arg0 set.Strings, arg1 error) *MockMachineAllSpacesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineAllSpacesCall) Do(f func(network.SubnetInfos) (set.Strings, error)) *MockMachineAllSpacesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineAllSpacesCall) DoAndReturn(f func(network.SubnetInfos) (set.Strings, error)) *MockMachineAllSpacesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Constraints mocks base method.
func (m *MockMachine) Constraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Constraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Constraints indicates an expected call of Constraints.
func (mr *MockMachineMockRecorder) Constraints() *MockMachineConstraintsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constraints", reflect.TypeOf((*MockMachine)(nil).Constraints))
	return &MockMachineConstraintsCall{Call: call}
}

// MockMachineConstraintsCall wrap *gomock.Call
type MockMachineConstraintsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineConstraintsCall) Return(arg0 constraints.Value, arg1 error) *MockMachineConstraintsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineConstraintsCall) Do(f func() (constraints.Value, error)) *MockMachineConstraintsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineConstraintsCall) DoAndReturn(f func() (constraints.Value, error)) *MockMachineConstraintsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ContainerType mocks base method.
func (m *MockMachine) ContainerType() instance.ContainerType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType.
func (mr *MockMachineMockRecorder) ContainerType() *MockMachineContainerTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockMachine)(nil).ContainerType))
	return &MockMachineContainerTypeCall{Call: call}
}

// MockMachineContainerTypeCall wrap *gomock.Call
type MockMachineContainerTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineContainerTypeCall) Return(arg0 instance.ContainerType) *MockMachineContainerTypeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineContainerTypeCall) Do(f func() instance.ContainerType) *MockMachineContainerTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineContainerTypeCall) DoAndReturn(f func() instance.ContainerType) *MockMachineContainerTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Id mocks base method.
func (m *MockMachine) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockMachineMockRecorder) Id() *MockMachineIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockMachine)(nil).Id))
	return &MockMachineIdCall{Call: call}
}

// MockMachineIdCall wrap *gomock.Call
type MockMachineIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineIdCall) Return(arg0 string) *MockMachineIdCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineIdCall) Do(f func() string) *MockMachineIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineIdCall) DoAndReturn(f func() string) *MockMachineIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InstanceId mocks base method.
func (m *MockMachine) InstanceId() (instance.Id, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceId")
	ret0, _ := ret[0].(instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceId indicates an expected call of InstanceId.
func (mr *MockMachineMockRecorder) InstanceId() *MockMachineInstanceIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceId", reflect.TypeOf((*MockMachine)(nil).InstanceId))
	return &MockMachineInstanceIdCall{Call: call}
}

// MockMachineInstanceIdCall wrap *gomock.Call
type MockMachineInstanceIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineInstanceIdCall) Return(arg0 instance.Id, arg1 error) *MockMachineInstanceIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineInstanceIdCall) Do(f func() (instance.Id, error)) *MockMachineInstanceIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineInstanceIdCall) DoAndReturn(f func() (instance.Id, error)) *MockMachineInstanceIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsManual mocks base method.
func (m *MockMachine) IsManual() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsManual")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsManual indicates an expected call of IsManual.
func (mr *MockMachineMockRecorder) IsManual() *MockMachineIsManualCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManual", reflect.TypeOf((*MockMachine)(nil).IsManual))
	return &MockMachineIsManualCall{Call: call}
}

// MockMachineIsManualCall wrap *gomock.Call
type MockMachineIsManualCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineIsManualCall) Return(arg0 bool, arg1 error) *MockMachineIsManualCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineIsManualCall) Do(f func() (bool, error)) *MockMachineIsManualCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineIsManualCall) DoAndReturn(f func() (bool, error)) *MockMachineIsManualCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MachineTag mocks base method.
func (m *MockMachine) MachineTag() names.MachineTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineTag")
	ret0, _ := ret[0].(names.MachineTag)
	return ret0
}

// MachineTag indicates an expected call of MachineTag.
func (mr *MockMachineMockRecorder) MachineTag() *MockMachineMachineTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineTag", reflect.TypeOf((*MockMachine)(nil).MachineTag))
	return &MockMachineMachineTagCall{Call: call}
}

// MockMachineMachineTagCall wrap *gomock.Call
type MockMachineMachineTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineMachineTagCall) Return(arg0 names.MachineTag) *MockMachineMachineTagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineMachineTagCall) Do(f func() names.MachineTag) *MockMachineMachineTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineMachineTagCall) DoAndReturn(f func() names.MachineTag) *MockMachineMachineTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Raw mocks base method.
func (m *MockMachine) Raw() *state.Machine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(*state.Machine)
	return ret0
}

// Raw indicates an expected call of Raw.
func (mr *MockMachineMockRecorder) Raw() *MockMachineRawCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockMachine)(nil).Raw))
	return &MockMachineRawCall{Call: call}
}

// MockMachineRawCall wrap *gomock.Call
type MockMachineRawCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineRawCall) Return(arg0 *state.Machine) *MockMachineRawCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineRawCall) Do(f func() *state.Machine) *MockMachineRawCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineRawCall) DoAndReturn(f func() *state.Machine) *MockMachineRawCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RemoveAllAddresses mocks base method.
func (m *MockMachine) RemoveAllAddresses() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllAddresses")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllAddresses indicates an expected call of RemoveAllAddresses.
func (mr *MockMachineMockRecorder) RemoveAllAddresses() *MockMachineRemoveAllAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllAddresses", reflect.TypeOf((*MockMachine)(nil).RemoveAllAddresses))
	return &MockMachineRemoveAllAddressesCall{Call: call}
}

// MockMachineRemoveAllAddressesCall wrap *gomock.Call
type MockMachineRemoveAllAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineRemoveAllAddressesCall) Return(arg0 error) *MockMachineRemoveAllAddressesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineRemoveAllAddressesCall) Do(f func() error) *MockMachineRemoveAllAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineRemoveAllAddressesCall) DoAndReturn(f func() error) *MockMachineRemoveAllAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetConstraints mocks base method.
func (m *MockMachine) SetConstraints(arg0 constraints.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConstraints", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConstraints indicates an expected call of SetConstraints.
func (mr *MockMachineMockRecorder) SetConstraints(arg0 any) *MockMachineSetConstraintsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConstraints", reflect.TypeOf((*MockMachine)(nil).SetConstraints), arg0)
	return &MockMachineSetConstraintsCall{Call: call}
}

// MockMachineSetConstraintsCall wrap *gomock.Call
type MockMachineSetConstraintsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineSetConstraintsCall) Return(arg0 error) *MockMachineSetConstraintsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineSetConstraintsCall) Do(f func(constraints.Value) error) *MockMachineSetConstraintsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineSetConstraintsCall) DoAndReturn(f func(constraints.Value) error) *MockMachineSetConstraintsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetDevicesAddresses mocks base method.
func (m *MockMachine) SetDevicesAddresses(arg0 ...state.LinkLayerDeviceAddress) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetDevicesAddresses", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDevicesAddresses indicates an expected call of SetDevicesAddresses.
func (mr *MockMachineMockRecorder) SetDevicesAddresses(arg0 ...any) *MockMachineSetDevicesAddressesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDevicesAddresses", reflect.TypeOf((*MockMachine)(nil).SetDevicesAddresses), arg0...)
	return &MockMachineSetDevicesAddressesCall{Call: call}
}

// MockMachineSetDevicesAddressesCall wrap *gomock.Call
type MockMachineSetDevicesAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineSetDevicesAddressesCall) Return(arg0 error) *MockMachineSetDevicesAddressesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineSetDevicesAddressesCall) Do(f func(...state.LinkLayerDeviceAddress) error) *MockMachineSetDevicesAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineSetDevicesAddressesCall) DoAndReturn(f func(...state.LinkLayerDeviceAddress) error) *MockMachineSetDevicesAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetLinkLayerDevices mocks base method.
func (m *MockMachine) SetLinkLayerDevices(arg0 ...state.LinkLayerDeviceArgs) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetLinkLayerDevices", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLinkLayerDevices indicates an expected call of SetLinkLayerDevices.
func (mr *MockMachineMockRecorder) SetLinkLayerDevices(arg0 ...any) *MockMachineSetLinkLayerDevicesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLinkLayerDevices", reflect.TypeOf((*MockMachine)(nil).SetLinkLayerDevices), arg0...)
	return &MockMachineSetLinkLayerDevicesCall{Call: call}
}

// MockMachineSetLinkLayerDevicesCall wrap *gomock.Call
type MockMachineSetLinkLayerDevicesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineSetLinkLayerDevicesCall) Return(arg0 error) *MockMachineSetLinkLayerDevicesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineSetLinkLayerDevicesCall) Do(f func(...state.LinkLayerDeviceArgs) error) *MockMachineSetLinkLayerDevicesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineSetLinkLayerDevicesCall) DoAndReturn(f func(...state.LinkLayerDeviceArgs) error) *MockMachineSetLinkLayerDevicesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Units mocks base method.
func (m *MockMachine) Units() ([]provisioner.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]provisioner.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units.
func (mr *MockMachineMockRecorder) Units() *MockMachineUnitsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockMachine)(nil).Units))
	return &MockMachineUnitsCall{Call: call}
}

// MockMachineUnitsCall wrap *gomock.Call
type MockMachineUnitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineUnitsCall) Return(arg0 []provisioner.Unit, arg1 error) *MockMachineUnitsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineUnitsCall) Do(f func() ([]provisioner.Unit, error)) *MockMachineUnitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineUnitsCall) DoAndReturn(f func() ([]provisioner.Unit, error)) *MockMachineUnitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockBridgePolicy is a mock of BridgePolicy interface.
type MockBridgePolicy struct {
	ctrl     *gomock.Controller
	recorder *MockBridgePolicyMockRecorder
}

// MockBridgePolicyMockRecorder is the mock recorder for MockBridgePolicy.
type MockBridgePolicyMockRecorder struct {
	mock *MockBridgePolicy
}

// NewMockBridgePolicy creates a new mock instance.
func NewMockBridgePolicy(ctrl *gomock.Controller) *MockBridgePolicy {
	mock := &MockBridgePolicy{ctrl: ctrl}
	mock.recorder = &MockBridgePolicyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBridgePolicy) EXPECT() *MockBridgePolicyMockRecorder {
	return m.recorder
}

// FindMissingBridgesForContainer mocks base method.
func (m *MockBridgePolicy) FindMissingBridgesForContainer(arg0 containerizer.Machine, arg1 containerizer.Container, arg2 network.SubnetInfos) ([]network0.DeviceToBridge, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMissingBridgesForContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].([]network0.DeviceToBridge)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindMissingBridgesForContainer indicates an expected call of FindMissingBridgesForContainer.
func (mr *MockBridgePolicyMockRecorder) FindMissingBridgesForContainer(arg0, arg1, arg2 any) *MockBridgePolicyFindMissingBridgesForContainerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMissingBridgesForContainer", reflect.TypeOf((*MockBridgePolicy)(nil).FindMissingBridgesForContainer), arg0, arg1, arg2)
	return &MockBridgePolicyFindMissingBridgesForContainerCall{Call: call}
}

// MockBridgePolicyFindMissingBridgesForContainerCall wrap *gomock.Call
type MockBridgePolicyFindMissingBridgesForContainerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBridgePolicyFindMissingBridgesForContainerCall) Return(arg0 []network0.DeviceToBridge, arg1 int, arg2 error) *MockBridgePolicyFindMissingBridgesForContainerCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBridgePolicyFindMissingBridgesForContainerCall) Do(f func(containerizer.Machine, containerizer.Container, network.SubnetInfos) ([]network0.DeviceToBridge, int, error)) *MockBridgePolicyFindMissingBridgesForContainerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBridgePolicyFindMissingBridgesForContainerCall) DoAndReturn(f func(containerizer.Machine, containerizer.Container, network.SubnetInfos) ([]network0.DeviceToBridge, int, error)) *MockBridgePolicyFindMissingBridgesForContainerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PopulateContainerLinkLayerDevices mocks base method.
func (m *MockBridgePolicy) PopulateContainerLinkLayerDevices(arg0 containerizer.Machine, arg1 containerizer.Container, arg2 bool) (network.InterfaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopulateContainerLinkLayerDevices", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.InterfaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PopulateContainerLinkLayerDevices indicates an expected call of PopulateContainerLinkLayerDevices.
func (mr *MockBridgePolicyMockRecorder) PopulateContainerLinkLayerDevices(arg0, arg1, arg2 any) *MockBridgePolicyPopulateContainerLinkLayerDevicesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateContainerLinkLayerDevices", reflect.TypeOf((*MockBridgePolicy)(nil).PopulateContainerLinkLayerDevices), arg0, arg1, arg2)
	return &MockBridgePolicyPopulateContainerLinkLayerDevicesCall{Call: call}
}

// MockBridgePolicyPopulateContainerLinkLayerDevicesCall wrap *gomock.Call
type MockBridgePolicyPopulateContainerLinkLayerDevicesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBridgePolicyPopulateContainerLinkLayerDevicesCall) Return(arg0 network.InterfaceInfos, arg1 error) *MockBridgePolicyPopulateContainerLinkLayerDevicesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBridgePolicyPopulateContainerLinkLayerDevicesCall) Do(f func(containerizer.Machine, containerizer.Container, bool) (network.InterfaceInfos, error)) *MockBridgePolicyPopulateContainerLinkLayerDevicesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBridgePolicyPopulateContainerLinkLayerDevicesCall) DoAndReturn(f func(containerizer.Machine, containerizer.Container, bool) (network.InterfaceInfos, error)) *MockBridgePolicyPopulateContainerLinkLayerDevicesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockUnit is a mock of Unit interface.
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit.
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance.
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// Application mocks base method.
func (m *MockUnit) Application() (provisioner.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application")
	ret0, _ := ret[0].(provisioner.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockUnitMockRecorder) Application() *MockUnitApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockUnit)(nil).Application))
	return &MockUnitApplicationCall{Call: call}
}

// MockUnitApplicationCall wrap *gomock.Call
type MockUnitApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUnitApplicationCall) Return(arg0 provisioner.Application, arg1 error) *MockUnitApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUnitApplicationCall) Do(f func() (provisioner.Application, error)) *MockUnitApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUnitApplicationCall) DoAndReturn(f func() (provisioner.Application, error)) *MockUnitApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockUnit) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockUnitMockRecorder) Name() *MockUnitNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockUnit)(nil).Name))
	return &MockUnitNameCall{Call: call}
}

// MockUnitNameCall wrap *gomock.Call
type MockUnitNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUnitNameCall) Return(arg0 string) *MockUnitNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUnitNameCall) Do(f func() string) *MockUnitNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUnitNameCall) DoAndReturn(f func() string) *MockUnitNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Charm mocks base method.
func (m *MockApplication) Charm() (provisioner.Charm, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm")
	ret0, _ := ret[0].(provisioner.Charm)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Charm indicates an expected call of Charm.
func (mr *MockApplicationMockRecorder) Charm() *MockApplicationCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockApplication)(nil).Charm))
	return &MockApplicationCharmCall{Call: call}
}

// MockApplicationCharmCall wrap *gomock.Call
type MockApplicationCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationCharmCall) Return(arg0 provisioner.Charm, arg1 bool, arg2 error) *MockApplicationCharmCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationCharmCall) Do(f func() (provisioner.Charm, bool, error)) *MockApplicationCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationCharmCall) DoAndReturn(f func() (provisioner.Charm, bool, error)) *MockApplicationCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Name mocks base method.
func (m *MockApplication) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockApplicationMockRecorder) Name() *MockApplicationNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockApplication)(nil).Name))
	return &MockApplicationNameCall{Call: call}
}

// MockApplicationNameCall wrap *gomock.Call
type MockApplicationNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationNameCall) Return(arg0 string) *MockApplicationNameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationNameCall) Do(f func() string) *MockApplicationNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationNameCall) DoAndReturn(f func() string) *MockApplicationNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockCharm is a mock of Charm interface.
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm.
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance.
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// LXDProfile mocks base method.
func (m *MockCharm) LXDProfile() *charm.LXDProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(*charm.LXDProfile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile.
func (mr *MockCharmMockRecorder) LXDProfile() *MockCharmLXDProfileCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockCharm)(nil).LXDProfile))
	return &MockCharmLXDProfileCall{Call: call}
}

// MockCharmLXDProfileCall wrap *gomock.Call
type MockCharmLXDProfileCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmLXDProfileCall) Return(arg0 *charm.LXDProfile) *MockCharmLXDProfileCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmLXDProfileCall) Do(f func() *charm.LXDProfile) *MockCharmLXDProfileCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmLXDProfileCall) DoAndReturn(f func() *charm.LXDProfile) *MockCharmLXDProfileCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Revision mocks base method.
func (m *MockCharm) Revision() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revision")
	ret0, _ := ret[0].(int)
	return ret0
}

// Revision indicates an expected call of Revision.
func (mr *MockCharmMockRecorder) Revision() *MockCharmRevisionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revision", reflect.TypeOf((*MockCharm)(nil).Revision))
	return &MockCharmRevisionCall{Call: call}
}

// MockCharmRevisionCall wrap *gomock.Call
type MockCharmRevisionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmRevisionCall) Return(arg0 int) *MockCharmRevisionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmRevisionCall) Do(f func() int) *MockCharmRevisionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmRevisionCall) DoAndReturn(f func() int) *MockCharmRevisionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
