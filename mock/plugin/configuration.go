// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/v2/plugin (interfaces: Configuration)

// Package plugin is a generated GoMock package.
package plugin

import (
	models "github.com/baetyl/baetyl-cloud/v2/models"
	v1 "github.com/baetyl/baetyl-go/v2/spec/v1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConfiguration is a mock of Configuration interface
type MockConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationMockRecorder
}

// MockConfigurationMockRecorder is the mock recorder for MockConfiguration
type MockConfigurationMockRecorder struct {
	mock *MockConfiguration
}

// NewMockConfiguration creates a new mock instance
func NewMockConfiguration(ctrl *gomock.Controller) *MockConfiguration {
	mock := &MockConfiguration{ctrl: ctrl}
	mock.recorder = &MockConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfiguration) EXPECT() *MockConfigurationMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockConfiguration) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConfigurationMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfiguration)(nil).Close))
}

// CreateConfig mocks base method
func (m *MockConfiguration) CreateConfig(arg0 string, arg1 *v1.Configuration) (*v1.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConfig", arg0, arg1)
	ret0, _ := ret[0].(*v1.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConfig indicates an expected call of CreateConfig
func (mr *MockConfigurationMockRecorder) CreateConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfig", reflect.TypeOf((*MockConfiguration)(nil).CreateConfig), arg0, arg1)
}

// DeleteConfig mocks base method
func (m *MockConfiguration) DeleteConfig(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConfig indicates an expected call of DeleteConfig
func (mr *MockConfigurationMockRecorder) DeleteConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfig", reflect.TypeOf((*MockConfiguration)(nil).DeleteConfig), arg0, arg1)
}

// GetConfig mocks base method
func (m *MockConfiguration) GetConfig(arg0, arg1, arg2 string) (*v1.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockConfigurationMockRecorder) GetConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockConfiguration)(nil).GetConfig), arg0, arg1, arg2)
}

// ListConfig mocks base method
func (m *MockConfiguration) ListConfig(arg0 string, arg1 *models.ListOptions) (*models.ConfigurationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConfig", arg0, arg1)
	ret0, _ := ret[0].(*models.ConfigurationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfig indicates an expected call of ListConfig
func (mr *MockConfigurationMockRecorder) ListConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfig", reflect.TypeOf((*MockConfiguration)(nil).ListConfig), arg0, arg1)
}

// UpdateConfig mocks base method
func (m *MockConfiguration) UpdateConfig(arg0 string, arg1 *v1.Configuration) (*v1.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfig", arg0, arg1)
	ret0, _ := ret[0].(*v1.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConfig indicates an expected call of UpdateConfig
func (mr *MockConfigurationMockRecorder) UpdateConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockConfiguration)(nil).UpdateConfig), arg0, arg1)
}
