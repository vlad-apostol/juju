// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/rbac/v1 (interfaces: RbacV1Interface,ClusterRoleBindingInterface,ClusterRoleInterface,RoleInterface,RoleBindingInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/rbac/v1"
	v12 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	rest "k8s.io/client-go/rest"
)

// MockRbacV1Interface is a mock of RbacV1Interface interface.
type MockRbacV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockRbacV1InterfaceMockRecorder
}

// MockRbacV1InterfaceMockRecorder is the mock recorder for MockRbacV1Interface.
type MockRbacV1InterfaceMockRecorder struct {
	mock *MockRbacV1Interface
}

// NewMockRbacV1Interface creates a new mock instance.
func NewMockRbacV1Interface(ctrl *gomock.Controller) *MockRbacV1Interface {
	mock := &MockRbacV1Interface{ctrl: ctrl}
	mock.recorder = &MockRbacV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRbacV1Interface) EXPECT() *MockRbacV1InterfaceMockRecorder {
	return m.recorder
}

// ClusterRoleBindings mocks base method.
func (m *MockRbacV1Interface) ClusterRoleBindings() v12.ClusterRoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterRoleBindings")
	ret0, _ := ret[0].(v12.ClusterRoleBindingInterface)
	return ret0
}

// ClusterRoleBindings indicates an expected call of ClusterRoleBindings.
func (mr *MockRbacV1InterfaceMockRecorder) ClusterRoleBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRoleBindings", reflect.TypeOf((*MockRbacV1Interface)(nil).ClusterRoleBindings))
}

// ClusterRoles mocks base method.
func (m *MockRbacV1Interface) ClusterRoles() v12.ClusterRoleInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterRoles")
	ret0, _ := ret[0].(v12.ClusterRoleInterface)
	return ret0
}

// ClusterRoles indicates an expected call of ClusterRoles.
func (mr *MockRbacV1InterfaceMockRecorder) ClusterRoles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterRoles", reflect.TypeOf((*MockRbacV1Interface)(nil).ClusterRoles))
}

// RESTClient mocks base method.
func (m *MockRbacV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockRbacV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockRbacV1Interface)(nil).RESTClient))
}

// RoleBindings mocks base method.
func (m *MockRbacV1Interface) RoleBindings(arg0 string) v12.RoleBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoleBindings", arg0)
	ret0, _ := ret[0].(v12.RoleBindingInterface)
	return ret0
}

// RoleBindings indicates an expected call of RoleBindings.
func (mr *MockRbacV1InterfaceMockRecorder) RoleBindings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoleBindings", reflect.TypeOf((*MockRbacV1Interface)(nil).RoleBindings), arg0)
}

// Roles mocks base method.
func (m *MockRbacV1Interface) Roles(arg0 string) v12.RoleInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Roles", arg0)
	ret0, _ := ret[0].(v12.RoleInterface)
	return ret0
}

// Roles indicates an expected call of Roles.
func (mr *MockRbacV1InterfaceMockRecorder) Roles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Roles", reflect.TypeOf((*MockRbacV1Interface)(nil).Roles), arg0)
}

// MockClusterRoleBindingInterface is a mock of ClusterRoleBindingInterface interface.
type MockClusterRoleBindingInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleBindingInterfaceMockRecorder
}

// MockClusterRoleBindingInterfaceMockRecorder is the mock recorder for MockClusterRoleBindingInterface.
type MockClusterRoleBindingInterfaceMockRecorder struct {
	mock *MockClusterRoleBindingInterface
}

// NewMockClusterRoleBindingInterface creates a new mock instance.
func NewMockClusterRoleBindingInterface(ctrl *gomock.Controller) *MockClusterRoleBindingInterface {
	mock := &MockClusterRoleBindingInterface{ctrl: ctrl}
	mock.recorder = &MockClusterRoleBindingInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRoleBindingInterface) EXPECT() *MockClusterRoleBindingInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockClusterRoleBindingInterface) Apply(arg0 context.Context, arg1 *v11.ClusterRoleBindingApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockClusterRoleBindingInterface) Create(arg0 context.Context, arg1 *v1.ClusterRoleBinding, arg2 v10.CreateOptions) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockClusterRoleBindingInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockClusterRoleBindingInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockClusterRoleBindingInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockClusterRoleBindingInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ClusterRoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ClusterRoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockClusterRoleBindingInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockClusterRoleBindingInterface) Update(arg0 context.Context, arg1 *v1.ClusterRoleBinding, arg2 v10.UpdateOptions) (*v1.ClusterRoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockClusterRoleBindingInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockClusterRoleBindingInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterRoleBindingInterface)(nil).Watch), arg0, arg1)
}

// MockClusterRoleInterface is a mock of ClusterRoleInterface interface.
type MockClusterRoleInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRoleInterfaceMockRecorder
}

// MockClusterRoleInterfaceMockRecorder is the mock recorder for MockClusterRoleInterface.
type MockClusterRoleInterfaceMockRecorder struct {
	mock *MockClusterRoleInterface
}

// NewMockClusterRoleInterface creates a new mock instance.
func NewMockClusterRoleInterface(ctrl *gomock.Controller) *MockClusterRoleInterface {
	mock := &MockClusterRoleInterface{ctrl: ctrl}
	mock.recorder = &MockClusterRoleInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterRoleInterface) EXPECT() *MockClusterRoleInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockClusterRoleInterface) Apply(arg0 context.Context, arg1 *v11.ClusterRoleApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockClusterRoleInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockClusterRoleInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockClusterRoleInterface) Create(arg0 context.Context, arg1 *v1.ClusterRole, arg2 v10.CreateOptions) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockClusterRoleInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterRoleInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockClusterRoleInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterRoleInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterRoleInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockClusterRoleInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockClusterRoleInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockClusterRoleInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockClusterRoleInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClusterRoleInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterRoleInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockClusterRoleInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ClusterRoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ClusterRoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClusterRoleInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterRoleInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockClusterRoleInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockClusterRoleInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockClusterRoleInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockClusterRoleInterface) Update(arg0 context.Context, arg1 *v1.ClusterRole, arg2 v10.UpdateOptions) (*v1.ClusterRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ClusterRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockClusterRoleInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterRoleInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockClusterRoleInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockClusterRoleInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterRoleInterface)(nil).Watch), arg0, arg1)
}

// MockRoleInterface is a mock of RoleInterface interface.
type MockRoleInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRoleInterfaceMockRecorder
}

// MockRoleInterfaceMockRecorder is the mock recorder for MockRoleInterface.
type MockRoleInterfaceMockRecorder struct {
	mock *MockRoleInterface
}

// NewMockRoleInterface creates a new mock instance.
func NewMockRoleInterface(ctrl *gomock.Controller) *MockRoleInterface {
	mock := &MockRoleInterface{ctrl: ctrl}
	mock.recorder = &MockRoleInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleInterface) EXPECT() *MockRoleInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockRoleInterface) Apply(arg0 context.Context, arg1 *v11.RoleApplyConfiguration, arg2 v10.ApplyOptions) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockRoleInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockRoleInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockRoleInterface) Create(arg0 context.Context, arg1 *v1.Role, arg2 v10.CreateOptions) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoleInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockRoleInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockRoleInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockRoleInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockRoleInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockRoleInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRoleInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRoleInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockRoleInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.RoleList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.RoleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoleInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockRoleInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.Role, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockRoleInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockRoleInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockRoleInterface) Update(arg0 context.Context, arg1 *v1.Role, arg2 v10.UpdateOptions) (*v1.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRoleInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockRoleInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockRoleInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRoleInterface)(nil).Watch), arg0, arg1)
}

// MockRoleBindingInterface is a mock of RoleBindingInterface interface.
type MockRoleBindingInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRoleBindingInterfaceMockRecorder
}

// MockRoleBindingInterfaceMockRecorder is the mock recorder for MockRoleBindingInterface.
type MockRoleBindingInterfaceMockRecorder struct {
	mock *MockRoleBindingInterface
}

// NewMockRoleBindingInterface creates a new mock instance.
func NewMockRoleBindingInterface(ctrl *gomock.Controller) *MockRoleBindingInterface {
	mock := &MockRoleBindingInterface{ctrl: ctrl}
	mock.recorder = &MockRoleBindingInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleBindingInterface) EXPECT() *MockRoleBindingInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockRoleBindingInterface) Apply(arg0 context.Context, arg1 *v11.RoleBindingApplyConfiguration, arg2 v10.ApplyOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockRoleBindingInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockRoleBindingInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockRoleBindingInterface) Create(arg0 context.Context, arg1 *v1.RoleBinding, arg2 v10.CreateOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoleBindingInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleBindingInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockRoleBindingInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleBindingInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleBindingInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockRoleBindingInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockRoleBindingInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockRoleBindingInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockRoleBindingInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRoleBindingInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRoleBindingInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockRoleBindingInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.RoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.RoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRoleBindingInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleBindingInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockRoleBindingInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockRoleBindingInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockRoleBindingInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockRoleBindingInterface) Update(arg0 context.Context, arg1 *v1.RoleBinding, arg2 v10.UpdateOptions) (*v1.RoleBinding, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.RoleBinding)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRoleBindingInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleBindingInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockRoleBindingInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockRoleBindingInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRoleBindingInterface)(nil).Watch), arg0, arg1)
}
