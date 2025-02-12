// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: NamespacesClient,TopicsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	servicebus "github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	gomock "github.com/golang/mock/gomock"
)

// MockNamespacesClient is a mock of NamespacesClient interface.
type MockNamespacesClient struct {
	ctrl     *gomock.Controller
	recorder *MockNamespacesClientMockRecorder
}

// MockNamespacesClientMockRecorder is the mock recorder for MockNamespacesClient.
type MockNamespacesClientMockRecorder struct {
	mock *MockNamespacesClient
}

// NewMockNamespacesClient creates a new mock instance.
func NewMockNamespacesClient(ctrl *gomock.Controller) *MockNamespacesClient {
	mock := &MockNamespacesClient{ctrl: ctrl}
	mock.recorder = &MockNamespacesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespacesClient) EXPECT() *MockNamespacesClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockNamespacesClient) List(arg0 context.Context) (servicebus.SBNamespaceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(servicebus.SBNamespaceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockNamespacesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNamespacesClient)(nil).List), arg0)
}

// MockTopicsClient is a mock of TopicsClient interface.
type MockTopicsClient struct {
	ctrl     *gomock.Controller
	recorder *MockTopicsClientMockRecorder
}

// MockTopicsClientMockRecorder is the mock recorder for MockTopicsClient.
type MockTopicsClientMockRecorder struct {
	mock *MockTopicsClient
}

// NewMockTopicsClient creates a new mock instance.
func NewMockTopicsClient(ctrl *gomock.Controller) *MockTopicsClient {
	mock := &MockTopicsClient{ctrl: ctrl}
	mock.recorder = &MockTopicsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopicsClient) EXPECT() *MockTopicsClientMockRecorder {
	return m.recorder
}

// ListAuthorizationRules mocks base method.
func (m *MockTopicsClient) ListAuthorizationRules(arg0 context.Context, arg1, arg2, arg3 string) (servicebus.SBAuthorizationRuleListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthorizationRules", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(servicebus.SBAuthorizationRuleListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthorizationRules indicates an expected call of ListAuthorizationRules.
func (mr *MockTopicsClientMockRecorder) ListAuthorizationRules(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthorizationRules", reflect.TypeOf((*MockTopicsClient)(nil).ListAuthorizationRules), arg0, arg1, arg2, arg3)
}

// ListByNamespace mocks base method.
func (m *MockTopicsClient) ListByNamespace(arg0 context.Context, arg1, arg2 string, arg3, arg4 *int32) (servicebus.SBTopicListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByNamespace", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(servicebus.SBTopicListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByNamespace indicates an expected call of ListByNamespace.
func (mr *MockTopicsClientMockRecorder) ListByNamespace(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByNamespace", reflect.TypeOf((*MockTopicsClient)(nil).ListByNamespace), arg0, arg1, arg2, arg3, arg4)
}

// ListKeys mocks base method.
func (m *MockTopicsClient) ListKeys(arg0 context.Context, arg1, arg2, arg3, arg4 string) (servicebus.AccessKeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(servicebus.AccessKeys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKeys indicates an expected call of ListKeys.
func (mr *MockTopicsClientMockRecorder) ListKeys(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockTopicsClient)(nil).ListKeys), arg0, arg1, arg2, arg3, arg4)
}
