// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/firewaller (interfaces: FirewallerAPI,RemoteRelationsAPI,CrossModelFirewallerFacadeCloser,EnvironFirewaller,EnvironModelFirewaller,EnvironInstances,EnvironInstance)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	api "github.com/juju/juju/api"
	firewaller "github.com/juju/juju/api/controller/firewaller"
	instance "github.com/juju/juju/core/instance"
	network "github.com/juju/juju/core/network"
	firewall "github.com/juju/juju/core/network/firewall"
	relation "github.com/juju/juju/core/relation"
	watcher "github.com/juju/juju/core/watcher"
	config "github.com/juju/juju/environs/config"
	context "github.com/juju/juju/environs/context"
	instances "github.com/juju/juju/environs/instances"
	params "github.com/juju/juju/rpc/params"
	firewaller0 "github.com/juju/juju/worker/firewaller"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
	macaroon "gopkg.in/macaroon.v2"
)

// MockFirewallerAPI is a mock of FirewallerAPI interface.
type MockFirewallerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallerAPIMockRecorder
}

// MockFirewallerAPIMockRecorder is the mock recorder for MockFirewallerAPI.
type MockFirewallerAPIMockRecorder struct {
	mock *MockFirewallerAPI
}

// NewMockFirewallerAPI creates a new mock instance.
func NewMockFirewallerAPI(ctrl *gomock.Controller) *MockFirewallerAPI {
	mock := &MockFirewallerAPI{ctrl: ctrl}
	mock.recorder = &MockFirewallerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirewallerAPI) EXPECT() *MockFirewallerAPIMockRecorder {
	return m.recorder
}

// AllSpaceInfos mocks base method.
func (m *MockFirewallerAPI) AllSpaceInfos() (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllSpaceInfos")
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllSpaceInfos indicates an expected call of AllSpaceInfos.
func (mr *MockFirewallerAPIMockRecorder) AllSpaceInfos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSpaceInfos", reflect.TypeOf((*MockFirewallerAPI)(nil).AllSpaceInfos))
}

// ControllerAPIInfoForModel mocks base method.
func (m *MockFirewallerAPI) ControllerAPIInfoForModel(arg0 string) (*api.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerAPIInfoForModel", arg0)
	ret0, _ := ret[0].(*api.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerAPIInfoForModel indicates an expected call of ControllerAPIInfoForModel.
func (mr *MockFirewallerAPIMockRecorder) ControllerAPIInfoForModel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerAPIInfoForModel", reflect.TypeOf((*MockFirewallerAPI)(nil).ControllerAPIInfoForModel), arg0)
}

// MacaroonForRelation mocks base method.
func (m *MockFirewallerAPI) MacaroonForRelation(arg0 string) (*macaroon.Macaroon, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MacaroonForRelation", arg0)
	ret0, _ := ret[0].(*macaroon.Macaroon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MacaroonForRelation indicates an expected call of MacaroonForRelation.
func (mr *MockFirewallerAPIMockRecorder) MacaroonForRelation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MacaroonForRelation", reflect.TypeOf((*MockFirewallerAPI)(nil).MacaroonForRelation), arg0)
}

// Machine mocks base method.
func (m *MockFirewallerAPI) Machine(arg0 names.MachineTag) (firewaller0.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(firewaller0.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockFirewallerAPIMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockFirewallerAPI)(nil).Machine), arg0)
}

// ModelConfig mocks base method.
func (m *MockFirewallerAPI) ModelConfig() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockFirewallerAPIMockRecorder) ModelConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockFirewallerAPI)(nil).ModelConfig))
}

// ModelFirewallRules mocks base method.
func (m *MockFirewallerAPI) ModelFirewallRules() (firewall.IngressRules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelFirewallRules")
	ret0, _ := ret[0].(firewall.IngressRules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelFirewallRules indicates an expected call of ModelFirewallRules.
func (mr *MockFirewallerAPIMockRecorder) ModelFirewallRules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelFirewallRules", reflect.TypeOf((*MockFirewallerAPI)(nil).ModelFirewallRules))
}

// Relation mocks base method.
func (m *MockFirewallerAPI) Relation(arg0 names.RelationTag) (*firewaller.Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relation", arg0)
	ret0, _ := ret[0].(*firewaller.Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relation indicates an expected call of Relation.
func (mr *MockFirewallerAPIMockRecorder) Relation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relation", reflect.TypeOf((*MockFirewallerAPI)(nil).Relation), arg0)
}

// SetRelationStatus mocks base method.
func (m *MockFirewallerAPI) SetRelationStatus(arg0 string, arg1 relation.Status, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRelationStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRelationStatus indicates an expected call of SetRelationStatus.
func (mr *MockFirewallerAPIMockRecorder) SetRelationStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRelationStatus", reflect.TypeOf((*MockFirewallerAPI)(nil).SetRelationStatus), arg0, arg1, arg2)
}

// Unit mocks base method.
func (m *MockFirewallerAPI) Unit(arg0 names.UnitTag) (firewaller0.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(firewaller0.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockFirewallerAPIMockRecorder) Unit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockFirewallerAPI)(nil).Unit), arg0)
}

// WatchEgressAddressesForRelation mocks base method.
func (m *MockFirewallerAPI) WatchEgressAddressesForRelation(arg0 names.RelationTag) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchEgressAddressesForRelation", arg0)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchEgressAddressesForRelation indicates an expected call of WatchEgressAddressesForRelation.
func (mr *MockFirewallerAPIMockRecorder) WatchEgressAddressesForRelation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchEgressAddressesForRelation", reflect.TypeOf((*MockFirewallerAPI)(nil).WatchEgressAddressesForRelation), arg0)
}

// WatchIngressAddressesForRelation mocks base method.
func (m *MockFirewallerAPI) WatchIngressAddressesForRelation(arg0 names.RelationTag) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchIngressAddressesForRelation", arg0)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchIngressAddressesForRelation indicates an expected call of WatchIngressAddressesForRelation.
func (mr *MockFirewallerAPIMockRecorder) WatchIngressAddressesForRelation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchIngressAddressesForRelation", reflect.TypeOf((*MockFirewallerAPI)(nil).WatchIngressAddressesForRelation), arg0)
}

// WatchModelFirewallRules mocks base method.
func (m *MockFirewallerAPI) WatchModelFirewallRules() (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchModelFirewallRules")
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchModelFirewallRules indicates an expected call of WatchModelFirewallRules.
func (mr *MockFirewallerAPIMockRecorder) WatchModelFirewallRules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchModelFirewallRules", reflect.TypeOf((*MockFirewallerAPI)(nil).WatchModelFirewallRules))
}

// WatchModelMachines mocks base method.
func (m *MockFirewallerAPI) WatchModelMachines() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchModelMachines")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchModelMachines indicates an expected call of WatchModelMachines.
func (mr *MockFirewallerAPIMockRecorder) WatchModelMachines() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchModelMachines", reflect.TypeOf((*MockFirewallerAPI)(nil).WatchModelMachines))
}

// WatchOpenedPorts mocks base method.
func (m *MockFirewallerAPI) WatchOpenedPorts() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOpenedPorts")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchOpenedPorts indicates an expected call of WatchOpenedPorts.
func (mr *MockFirewallerAPIMockRecorder) WatchOpenedPorts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOpenedPorts", reflect.TypeOf((*MockFirewallerAPI)(nil).WatchOpenedPorts))
}

// WatchSubnets mocks base method.
func (m *MockFirewallerAPI) WatchSubnets() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchSubnets")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchSubnets indicates an expected call of WatchSubnets.
func (mr *MockFirewallerAPIMockRecorder) WatchSubnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchSubnets", reflect.TypeOf((*MockFirewallerAPI)(nil).WatchSubnets))
}

// MockRemoteRelationsAPI is a mock of RemoteRelationsAPI interface.
type MockRemoteRelationsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteRelationsAPIMockRecorder
}

// MockRemoteRelationsAPIMockRecorder is the mock recorder for MockRemoteRelationsAPI.
type MockRemoteRelationsAPIMockRecorder struct {
	mock *MockRemoteRelationsAPI
}

// NewMockRemoteRelationsAPI creates a new mock instance.
func NewMockRemoteRelationsAPI(ctrl *gomock.Controller) *MockRemoteRelationsAPI {
	mock := &MockRemoteRelationsAPI{ctrl: ctrl}
	mock.recorder = &MockRemoteRelationsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteRelationsAPI) EXPECT() *MockRemoteRelationsAPIMockRecorder {
	return m.recorder
}

// GetToken mocks base method.
func (m *MockRemoteRelationsAPI) GetToken(arg0 names.Tag) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockRemoteRelationsAPIMockRecorder) GetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockRemoteRelationsAPI)(nil).GetToken), arg0)
}

// Relations mocks base method.
func (m *MockRemoteRelationsAPI) Relations(arg0 []string) ([]params.RemoteRelationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relations", arg0)
	ret0, _ := ret[0].([]params.RemoteRelationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relations indicates an expected call of Relations.
func (mr *MockRemoteRelationsAPIMockRecorder) Relations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relations", reflect.TypeOf((*MockRemoteRelationsAPI)(nil).Relations), arg0)
}

// RemoteApplications mocks base method.
func (m *MockRemoteRelationsAPI) RemoteApplications(arg0 []string) ([]params.RemoteApplicationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteApplications", arg0)
	ret0, _ := ret[0].([]params.RemoteApplicationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteApplications indicates an expected call of RemoteApplications.
func (mr *MockRemoteRelationsAPIMockRecorder) RemoteApplications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteApplications", reflect.TypeOf((*MockRemoteRelationsAPI)(nil).RemoteApplications), arg0)
}

// WatchRemoteRelations mocks base method.
func (m *MockRemoteRelationsAPI) WatchRemoteRelations() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchRemoteRelations")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchRemoteRelations indicates an expected call of WatchRemoteRelations.
func (mr *MockRemoteRelationsAPIMockRecorder) WatchRemoteRelations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchRemoteRelations", reflect.TypeOf((*MockRemoteRelationsAPI)(nil).WatchRemoteRelations))
}

// MockCrossModelFirewallerFacadeCloser is a mock of CrossModelFirewallerFacadeCloser interface.
type MockCrossModelFirewallerFacadeCloser struct {
	ctrl     *gomock.Controller
	recorder *MockCrossModelFirewallerFacadeCloserMockRecorder
}

// MockCrossModelFirewallerFacadeCloserMockRecorder is the mock recorder for MockCrossModelFirewallerFacadeCloser.
type MockCrossModelFirewallerFacadeCloserMockRecorder struct {
	mock *MockCrossModelFirewallerFacadeCloser
}

// NewMockCrossModelFirewallerFacadeCloser creates a new mock instance.
func NewMockCrossModelFirewallerFacadeCloser(ctrl *gomock.Controller) *MockCrossModelFirewallerFacadeCloser {
	mock := &MockCrossModelFirewallerFacadeCloser{ctrl: ctrl}
	mock.recorder = &MockCrossModelFirewallerFacadeCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrossModelFirewallerFacadeCloser) EXPECT() *MockCrossModelFirewallerFacadeCloserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCrossModelFirewallerFacadeCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCrossModelFirewallerFacadeCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCrossModelFirewallerFacadeCloser)(nil).Close))
}

// PublishIngressNetworkChange mocks base method.
func (m *MockCrossModelFirewallerFacadeCloser) PublishIngressNetworkChange(arg0 params.IngressNetworksChangeEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishIngressNetworkChange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishIngressNetworkChange indicates an expected call of PublishIngressNetworkChange.
func (mr *MockCrossModelFirewallerFacadeCloserMockRecorder) PublishIngressNetworkChange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishIngressNetworkChange", reflect.TypeOf((*MockCrossModelFirewallerFacadeCloser)(nil).PublishIngressNetworkChange), arg0)
}

// WatchEgressAddressesForRelation mocks base method.
func (m *MockCrossModelFirewallerFacadeCloser) WatchEgressAddressesForRelation(arg0 params.RemoteEntityArg) (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchEgressAddressesForRelation", arg0)
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchEgressAddressesForRelation indicates an expected call of WatchEgressAddressesForRelation.
func (mr *MockCrossModelFirewallerFacadeCloserMockRecorder) WatchEgressAddressesForRelation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchEgressAddressesForRelation", reflect.TypeOf((*MockCrossModelFirewallerFacadeCloser)(nil).WatchEgressAddressesForRelation), arg0)
}

// MockEnvironFirewaller is a mock of EnvironFirewaller interface.
type MockEnvironFirewaller struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironFirewallerMockRecorder
}

// MockEnvironFirewallerMockRecorder is the mock recorder for MockEnvironFirewaller.
type MockEnvironFirewallerMockRecorder struct {
	mock *MockEnvironFirewaller
}

// NewMockEnvironFirewaller creates a new mock instance.
func NewMockEnvironFirewaller(ctrl *gomock.Controller) *MockEnvironFirewaller {
	mock := &MockEnvironFirewaller{ctrl: ctrl}
	mock.recorder = &MockEnvironFirewallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironFirewaller) EXPECT() *MockEnvironFirewallerMockRecorder {
	return m.recorder
}

// ClosePorts mocks base method.
func (m *MockEnvironFirewaller) ClosePorts(arg0 context.ProviderCallContext, arg1 firewall.IngressRules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClosePorts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClosePorts indicates an expected call of ClosePorts.
func (mr *MockEnvironFirewallerMockRecorder) ClosePorts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosePorts", reflect.TypeOf((*MockEnvironFirewaller)(nil).ClosePorts), arg0, arg1)
}

// IngressRules mocks base method.
func (m *MockEnvironFirewaller) IngressRules(arg0 context.ProviderCallContext) (firewall.IngressRules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngressRules", arg0)
	ret0, _ := ret[0].(firewall.IngressRules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngressRules indicates an expected call of IngressRules.
func (mr *MockEnvironFirewallerMockRecorder) IngressRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngressRules", reflect.TypeOf((*MockEnvironFirewaller)(nil).IngressRules), arg0)
}

// OpenPorts mocks base method.
func (m *MockEnvironFirewaller) OpenPorts(arg0 context.ProviderCallContext, arg1 firewall.IngressRules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenPorts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenPorts indicates an expected call of OpenPorts.
func (mr *MockEnvironFirewallerMockRecorder) OpenPorts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPorts", reflect.TypeOf((*MockEnvironFirewaller)(nil).OpenPorts), arg0, arg1)
}

// MockEnvironModelFirewaller is a mock of EnvironModelFirewaller interface.
type MockEnvironModelFirewaller struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironModelFirewallerMockRecorder
}

// MockEnvironModelFirewallerMockRecorder is the mock recorder for MockEnvironModelFirewaller.
type MockEnvironModelFirewallerMockRecorder struct {
	mock *MockEnvironModelFirewaller
}

// NewMockEnvironModelFirewaller creates a new mock instance.
func NewMockEnvironModelFirewaller(ctrl *gomock.Controller) *MockEnvironModelFirewaller {
	mock := &MockEnvironModelFirewaller{ctrl: ctrl}
	mock.recorder = &MockEnvironModelFirewallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironModelFirewaller) EXPECT() *MockEnvironModelFirewallerMockRecorder {
	return m.recorder
}

// CloseModelPorts mocks base method.
func (m *MockEnvironModelFirewaller) CloseModelPorts(arg0 context.ProviderCallContext, arg1 firewall.IngressRules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseModelPorts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseModelPorts indicates an expected call of CloseModelPorts.
func (mr *MockEnvironModelFirewallerMockRecorder) CloseModelPorts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseModelPorts", reflect.TypeOf((*MockEnvironModelFirewaller)(nil).CloseModelPorts), arg0, arg1)
}

// ModelIngressRules mocks base method.
func (m *MockEnvironModelFirewaller) ModelIngressRules(arg0 context.ProviderCallContext) (firewall.IngressRules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelIngressRules", arg0)
	ret0, _ := ret[0].(firewall.IngressRules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelIngressRules indicates an expected call of ModelIngressRules.
func (mr *MockEnvironModelFirewallerMockRecorder) ModelIngressRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelIngressRules", reflect.TypeOf((*MockEnvironModelFirewaller)(nil).ModelIngressRules), arg0)
}

// OpenModelPorts mocks base method.
func (m *MockEnvironModelFirewaller) OpenModelPorts(arg0 context.ProviderCallContext, arg1 firewall.IngressRules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenModelPorts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenModelPorts indicates an expected call of OpenModelPorts.
func (mr *MockEnvironModelFirewallerMockRecorder) OpenModelPorts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenModelPorts", reflect.TypeOf((*MockEnvironModelFirewaller)(nil).OpenModelPorts), arg0, arg1)
}

// MockEnvironInstances is a mock of EnvironInstances interface.
type MockEnvironInstances struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironInstancesMockRecorder
}

// MockEnvironInstancesMockRecorder is the mock recorder for MockEnvironInstances.
type MockEnvironInstancesMockRecorder struct {
	mock *MockEnvironInstances
}

// NewMockEnvironInstances creates a new mock instance.
func NewMockEnvironInstances(ctrl *gomock.Controller) *MockEnvironInstances {
	mock := &MockEnvironInstances{ctrl: ctrl}
	mock.recorder = &MockEnvironInstancesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironInstances) EXPECT() *MockEnvironInstancesMockRecorder {
	return m.recorder
}

// Instances mocks base method.
func (m *MockEnvironInstances) Instances(arg0 context.ProviderCallContext, arg1 []instance.Id) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instances", arg0, arg1)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Instances indicates an expected call of Instances.
func (mr *MockEnvironInstancesMockRecorder) Instances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instances", reflect.TypeOf((*MockEnvironInstances)(nil).Instances), arg0, arg1)
}

// MockEnvironInstance is a mock of EnvironInstance interface.
type MockEnvironInstance struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironInstanceMockRecorder
}

// MockEnvironInstanceMockRecorder is the mock recorder for MockEnvironInstance.
type MockEnvironInstanceMockRecorder struct {
	mock *MockEnvironInstance
}

// NewMockEnvironInstance creates a new mock instance.
func NewMockEnvironInstance(ctrl *gomock.Controller) *MockEnvironInstance {
	mock := &MockEnvironInstance{ctrl: ctrl}
	mock.recorder = &MockEnvironInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEnvironInstance) EXPECT() *MockEnvironInstanceMockRecorder {
	return m.recorder
}

// Addresses mocks base method.
func (m *MockEnvironInstance) Addresses(arg0 context.ProviderCallContext) (network.ProviderAddresses, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses", arg0)
	ret0, _ := ret[0].(network.ProviderAddresses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Addresses indicates an expected call of Addresses.
func (mr *MockEnvironInstanceMockRecorder) Addresses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockEnvironInstance)(nil).Addresses), arg0)
}

// ClosePorts mocks base method.
func (m *MockEnvironInstance) ClosePorts(arg0 context.ProviderCallContext, arg1 string, arg2 firewall.IngressRules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClosePorts", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClosePorts indicates an expected call of ClosePorts.
func (mr *MockEnvironInstanceMockRecorder) ClosePorts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosePorts", reflect.TypeOf((*MockEnvironInstance)(nil).ClosePorts), arg0, arg1, arg2)
}

// Id mocks base method.
func (m *MockEnvironInstance) Id() instance.Id {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(instance.Id)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockEnvironInstanceMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockEnvironInstance)(nil).Id))
}

// IngressRules mocks base method.
func (m *MockEnvironInstance) IngressRules(arg0 context.ProviderCallContext, arg1 string) (firewall.IngressRules, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IngressRules", arg0, arg1)
	ret0, _ := ret[0].(firewall.IngressRules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IngressRules indicates an expected call of IngressRules.
func (mr *MockEnvironInstanceMockRecorder) IngressRules(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IngressRules", reflect.TypeOf((*MockEnvironInstance)(nil).IngressRules), arg0, arg1)
}

// OpenPorts mocks base method.
func (m *MockEnvironInstance) OpenPorts(arg0 context.ProviderCallContext, arg1 string, arg2 firewall.IngressRules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenPorts", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenPorts indicates an expected call of OpenPorts.
func (mr *MockEnvironInstanceMockRecorder) OpenPorts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenPorts", reflect.TypeOf((*MockEnvironInstance)(nil).OpenPorts), arg0, arg1, arg2)
}

// Status mocks base method.
func (m *MockEnvironInstance) Status(arg0 context.ProviderCallContext) instance.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(instance.Status)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockEnvironInstanceMockRecorder) Status(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockEnvironInstance)(nil).Status), arg0)
}
