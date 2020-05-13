// Code generated by MockGen. DO NOT EDIT.
// Source: appconf.go

// Package appconf is a generated GoMock package.
package appconf

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/oam-dev/cloud-provider/alibabacloud/ros/apis/ros.alibabacloud.com/v1alpha1"
	v1alpha10 "github.com/oam-dev/oam-go-sdk/apis/core.oam.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
)

// MockConfigurationInterface is a mock of AppConfInterface interface
type MockConfigurationInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationInterfaceMockRecorder
}

// MockConfigurationInterfaceMockRecorder is the mock recorder for MockConfigurationInterface
type MockConfigurationInterfaceMockRecorder struct {
	mock *MockConfigurationInterface
}

// NewMockConfigurationInterface creates a new mock instance
func NewMockConfigurationInterface(ctrl *gomock.Controller) *MockConfigurationInterface {
	mock := &MockConfigurationInterface{ctrl: ctrl}
	mock.recorder = &MockConfigurationInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigurationInterface) EXPECT() *MockConfigurationInterfaceMockRecorder {
	return m.recorder
}

// ToOamApplicationConfiguration mocks base method
func (m *MockConfigurationInterface) ToOamApplicationConfiguration() *v1alpha10.ApplicationConfiguration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToOamApplicationConfiguration")
	ret0, _ := ret[0].(*v1alpha10.ApplicationConfiguration)
	return ret0
}

// ToOamApplicationConfiguration indicates an expected call of ToOamApplicationConfiguration
func (mr *MockConfigurationInterfaceMockRecorder) ToOamApplicationConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToOamApplicationConfiguration", reflect.TypeOf((*MockConfigurationInterface)(nil).ToOamApplicationConfiguration))
}

// ToRosStack mocks base method
func (m *MockConfigurationInterface) ToRosStack() *v1alpha1.RosStack {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToRosStack")
	ret0, _ := ret[0].(*v1alpha1.RosStack)
	return ret0
}

// ToRosStack indicates an expected call of ToRosStack
func (mr *MockConfigurationInterfaceMockRecorder) ToRosStack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToRosStack", reflect.TypeOf((*MockConfigurationInterface)(nil).ToRosStack))
}

// ToObject mocks base method
func (m *MockConfigurationInterface) ToObject() v1.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToObject")
	ret0, _ := ret[0].(v1.Object)
	return ret0
}

// ToObject indicates an expected call of ToObject
func (mr *MockConfigurationInterfaceMockRecorder) ToObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToObject", reflect.TypeOf((*MockConfigurationInterface)(nil).ToObject))
}

// Update mocks base method
func (m *MockConfigurationInterface) Update(c *Context, appConf AppConfInterface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, appConf)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockConfigurationInterfaceMockRecorder) Update(c, appConf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockConfigurationInterface)(nil).Update), c, appConf)
}

// UpdateStatus mocks base method
func (m *MockConfigurationInterface) UpdateStatus(c *Context, phase v1alpha10.ApplicationPhase, type_ v1alpha10.ApplicationConditionType, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", c, phase, type_, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockConfigurationInterfaceMockRecorder) UpdateStatus(c, phase, type_, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockConfigurationInterface)(nil).UpdateStatus), c, phase, type_, message)
}

// GetScopes mocks base method
func (m *MockConfigurationInterface) GetScopes() []v1alpha10.ScopeBinding {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScopes")
	ret0, _ := ret[0].([]v1alpha10.ScopeBinding)
	return ret0
}

// GetScopes indicates an expected call of GetScopes
func (mr *MockConfigurationInterfaceMockRecorder) GetScopes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScopes", reflect.TypeOf((*MockConfigurationInterface)(nil).GetScopes))
}

// GetName mocks base method
func (m *MockConfigurationInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockConfigurationInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockConfigurationInterface)(nil).GetName))
}

// GetNamespace mocks base method
func (m *MockConfigurationInterface) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockConfigurationInterfaceMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockConfigurationInterface)(nil).GetNamespace))
}

// GetObjectMeta mocks base method
func (m *MockConfigurationInterface) GetObjectMeta() v1.ObjectMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectMeta")
	ret0, _ := ret[0].(v1.ObjectMeta)
	return ret0
}

// GetObjectMeta indicates an expected call of GetObjectMeta
func (mr *MockConfigurationInterfaceMockRecorder) GetObjectMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMeta", reflect.TypeOf((*MockConfigurationInterface)(nil).GetObjectMeta))
}