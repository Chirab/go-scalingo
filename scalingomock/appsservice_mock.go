// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo/v4 (interfaces: AppsService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	http "net/http"
	reflect "reflect"

	scalingo "github.com/Scalingo/go-scalingo/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockAppsService is a mock of AppsService interface
type MockAppsService struct {
	ctrl     *gomock.Controller
	recorder *MockAppsServiceMockRecorder
}

// MockAppsServiceMockRecorder is the mock recorder for MockAppsService
type MockAppsServiceMockRecorder struct {
	mock *MockAppsService
}

// NewMockAppsService creates a new mock instance
func NewMockAppsService(ctrl *gomock.Controller) *MockAppsService {
	mock := &MockAppsService{ctrl: ctrl}
	mock.recorder = &MockAppsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppsService) EXPECT() *MockAppsServiceMockRecorder {
	return m.recorder
}

// AppsCreate mocks base method
func (m *MockAppsService) AppsCreate(arg0 scalingo.AppsCreateOpts) (*scalingo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsCreate", arg0)
	ret0, _ := ret[0].(*scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsCreate indicates an expected call of AppsCreate
func (mr *MockAppsServiceMockRecorder) AppsCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsCreate", reflect.TypeOf((*MockAppsService)(nil).AppsCreate), arg0)
}

// AppsDestroy mocks base method
func (m *MockAppsService) AppsDestroy(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppsDestroy indicates an expected call of AppsDestroy
func (mr *MockAppsServiceMockRecorder) AppsDestroy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsDestroy", reflect.TypeOf((*MockAppsService)(nil).AppsDestroy), arg0, arg1)
}

// AppsForceHTTPS mocks base method
func (m *MockAppsService) AppsForceHTTPS(arg0 string, arg1 bool) (*scalingo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsForceHTTPS", arg0, arg1)
	ret0, _ := ret[0].(*scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsForceHTTPS indicates an expected call of AppsForceHTTPS
func (mr *MockAppsServiceMockRecorder) AppsForceHTTPS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsForceHTTPS", reflect.TypeOf((*MockAppsService)(nil).AppsForceHTTPS), arg0, arg1)
}

// AppsList mocks base method
func (m *MockAppsService) AppsList() ([]*scalingo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsList")
	ret0, _ := ret[0].([]*scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsList indicates an expected call of AppsList
func (mr *MockAppsServiceMockRecorder) AppsList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsList", reflect.TypeOf((*MockAppsService)(nil).AppsList))
}

// AppsPs mocks base method
func (m *MockAppsService) AppsPs(arg0 string) ([]scalingo.ContainerType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsPs", arg0)
	ret0, _ := ret[0].([]scalingo.ContainerType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsPs indicates an expected call of AppsPs
func (mr *MockAppsServiceMockRecorder) AppsPs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsPs", reflect.TypeOf((*MockAppsService)(nil).AppsPs), arg0)
}

// AppsRename mocks base method
func (m *MockAppsService) AppsRename(arg0, arg1 string) (*scalingo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsRename", arg0, arg1)
	ret0, _ := ret[0].(*scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsRename indicates an expected call of AppsRename
func (mr *MockAppsServiceMockRecorder) AppsRename(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsRename", reflect.TypeOf((*MockAppsService)(nil).AppsRename), arg0, arg1)
}

// AppsRestart mocks base method
func (m *MockAppsService) AppsRestart(arg0 string, arg1 *scalingo.AppsRestartParams) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsRestart", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsRestart indicates an expected call of AppsRestart
func (mr *MockAppsServiceMockRecorder) AppsRestart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsRestart", reflect.TypeOf((*MockAppsService)(nil).AppsRestart), arg0, arg1)
}

// AppsScale mocks base method
func (m *MockAppsService) AppsScale(arg0 string, arg1 *scalingo.AppsScaleParams) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsScale", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsScale indicates an expected call of AppsScale
func (mr *MockAppsServiceMockRecorder) AppsScale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsScale", reflect.TypeOf((*MockAppsService)(nil).AppsScale), arg0, arg1)
}

// AppsSetStack mocks base method
func (m *MockAppsService) AppsSetStack(arg0, arg1 string) (*scalingo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsSetStack", arg0, arg1)
	ret0, _ := ret[0].(*scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsSetStack indicates an expected call of AppsSetStack
func (mr *MockAppsServiceMockRecorder) AppsSetStack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsSetStack", reflect.TypeOf((*MockAppsService)(nil).AppsSetStack), arg0, arg1)
}

// AppsShow mocks base method
func (m *MockAppsService) AppsShow(arg0 string) (*scalingo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsShow", arg0)
	ret0, _ := ret[0].(*scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsShow indicates an expected call of AppsShow
func (mr *MockAppsServiceMockRecorder) AppsShow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsShow", reflect.TypeOf((*MockAppsService)(nil).AppsShow), arg0)
}

// AppsStats mocks base method
func (m *MockAppsService) AppsStats(arg0 string) (*scalingo.AppStatsRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsStats", arg0)
	ret0, _ := ret[0].(*scalingo.AppStatsRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsStats indicates an expected call of AppsStats
func (mr *MockAppsServiceMockRecorder) AppsStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsStats", reflect.TypeOf((*MockAppsService)(nil).AppsStats), arg0)
}

// AppsStickySession mocks base method
func (m *MockAppsService) AppsStickySession(arg0 string, arg1 bool) (*scalingo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsStickySession", arg0, arg1)
	ret0, _ := ret[0].(*scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsStickySession indicates an expected call of AppsStickySession
func (mr *MockAppsServiceMockRecorder) AppsStickySession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsStickySession", reflect.TypeOf((*MockAppsService)(nil).AppsStickySession), arg0, arg1)
}

// AppsTransfer mocks base method
func (m *MockAppsService) AppsTransfer(arg0, arg1 string) (*scalingo.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsTransfer", arg0, arg1)
	ret0, _ := ret[0].(*scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsTransfer indicates an expected call of AppsTransfer
func (mr *MockAppsServiceMockRecorder) AppsTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsTransfer", reflect.TypeOf((*MockAppsService)(nil).AppsTransfer), arg0, arg1)
}
