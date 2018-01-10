// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: Client)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	websocket "golang.org/x/net/websocket"
	http "net/http"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddonDestroy mocks base method
func (m *MockClient) AddonDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AddonDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddonDestroy indicates an expected call of AddonDestroy
func (mr *MockClientMockRecorder) AddonDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonDestroy", reflect.TypeOf((*MockClient)(nil).AddonDestroy), arg0, arg1)
}

// AddonProviderPlansList mocks base method
func (m *MockClient) AddonProviderPlansList(arg0 string) ([]*go_scalingo.Plan, error) {
	ret := m.ctrl.Call(m, "AddonProviderPlansList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProviderPlansList indicates an expected call of AddonProviderPlansList
func (mr *MockClientMockRecorder) AddonProviderPlansList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProviderPlansList", reflect.TypeOf((*MockClient)(nil).AddonProviderPlansList), arg0)
}

// AddonProvidersList mocks base method
func (m *MockClient) AddonProvidersList() ([]*go_scalingo.AddonProvider, error) {
	ret := m.ctrl.Call(m, "AddonProvidersList")
	ret0, _ := ret[0].([]*go_scalingo.AddonProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProvidersList indicates an expected call of AddonProvidersList
func (mr *MockClientMockRecorder) AddonProvidersList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProvidersList", reflect.TypeOf((*MockClient)(nil).AddonProvidersList))
}

// AddonProvision mocks base method
func (m *MockClient) AddonProvision(arg0, arg1, arg2 string) (go_scalingo.AddonRes, error) {
	ret := m.ctrl.Call(m, "AddonProvision", arg0, arg1, arg2)
	ret0, _ := ret[0].(go_scalingo.AddonRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProvision indicates an expected call of AddonProvision
func (mr *MockClientMockRecorder) AddonProvision(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProvision", reflect.TypeOf((*MockClient)(nil).AddonProvision), arg0, arg1, arg2)
}

// AddonUpgrade mocks base method
func (m *MockClient) AddonUpgrade(arg0, arg1, arg2 string) (go_scalingo.AddonRes, error) {
	ret := m.ctrl.Call(m, "AddonUpgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(go_scalingo.AddonRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonUpgrade indicates an expected call of AddonUpgrade
func (mr *MockClientMockRecorder) AddonUpgrade(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonUpgrade", reflect.TypeOf((*MockClient)(nil).AddonUpgrade), arg0, arg1, arg2)
}

// AddonsList mocks base method
func (m *MockClient) AddonsList(arg0 string) ([]*go_scalingo.Addon, error) {
	ret := m.ctrl.Call(m, "AddonsList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Addon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonsList indicates an expected call of AddonsList
func (mr *MockClientMockRecorder) AddonsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonsList", reflect.TypeOf((*MockClient)(nil).AddonsList), arg0)
}

// AppsCreate mocks base method
func (m *MockClient) AppsCreate(arg0 go_scalingo.AppsCreateOpts) (*go_scalingo.App, error) {
	ret := m.ctrl.Call(m, "AppsCreate", arg0)
	ret0, _ := ret[0].(*go_scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsCreate indicates an expected call of AppsCreate
func (mr *MockClientMockRecorder) AppsCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsCreate", reflect.TypeOf((*MockClient)(nil).AppsCreate), arg0)
}

// AppsDestroy mocks base method
func (m *MockClient) AppsDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AppsDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppsDestroy indicates an expected call of AppsDestroy
func (mr *MockClientMockRecorder) AppsDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsDestroy", reflect.TypeOf((*MockClient)(nil).AppsDestroy), arg0, arg1)
}

// AppsList mocks base method
func (m *MockClient) AppsList() ([]*go_scalingo.App, error) {
	ret := m.ctrl.Call(m, "AppsList")
	ret0, _ := ret[0].([]*go_scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsList indicates an expected call of AppsList
func (mr *MockClientMockRecorder) AppsList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsList", reflect.TypeOf((*MockClient)(nil).AppsList))
}

// AppsPs mocks base method
func (m *MockClient) AppsPs(arg0 string) ([]go_scalingo.ContainerType, error) {
	ret := m.ctrl.Call(m, "AppsPs", arg0)
	ret0, _ := ret[0].([]go_scalingo.ContainerType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsPs indicates an expected call of AppsPs
func (mr *MockClientMockRecorder) AppsPs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsPs", reflect.TypeOf((*MockClient)(nil).AppsPs), arg0)
}

// AppsRestart mocks base method
func (m *MockClient) AppsRestart(arg0 string, arg1 *go_scalingo.AppsRestartParams) (*http.Response, error) {
	ret := m.ctrl.Call(m, "AppsRestart", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsRestart indicates an expected call of AppsRestart
func (mr *MockClientMockRecorder) AppsRestart(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsRestart", reflect.TypeOf((*MockClient)(nil).AppsRestart), arg0, arg1)
}

// AppsScale mocks base method
func (m *MockClient) AppsScale(arg0 string, arg1 *go_scalingo.AppsScaleParams) (*http.Response, error) {
	ret := m.ctrl.Call(m, "AppsScale", arg0, arg1)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsScale indicates an expected call of AppsScale
func (mr *MockClientMockRecorder) AppsScale(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsScale", reflect.TypeOf((*MockClient)(nil).AppsScale), arg0, arg1)
}

// AppsShow mocks base method
func (m *MockClient) AppsShow(arg0 string) (*go_scalingo.App, error) {
	ret := m.ctrl.Call(m, "AppsShow", arg0)
	ret0, _ := ret[0].(*go_scalingo.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsShow indicates an expected call of AppsShow
func (mr *MockClientMockRecorder) AppsShow(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsShow", reflect.TypeOf((*MockClient)(nil).AppsShow), arg0)
}

// AppsStats mocks base method
func (m *MockClient) AppsStats(arg0 string) (*go_scalingo.AppStatsRes, error) {
	ret := m.ctrl.Call(m, "AppsStats", arg0)
	ret0, _ := ret[0].(*go_scalingo.AppStatsRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppsStats indicates an expected call of AppsStats
func (mr *MockClientMockRecorder) AppsStats(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsStats", reflect.TypeOf((*MockClient)(nil).AppsStats), arg0)
}

// CollaboratorAdd mocks base method
func (m *MockClient) CollaboratorAdd(arg0, arg1 string) (go_scalingo.Collaborator, error) {
	ret := m.ctrl.Call(m, "CollaboratorAdd", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Collaborator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollaboratorAdd indicates an expected call of CollaboratorAdd
func (mr *MockClientMockRecorder) CollaboratorAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollaboratorAdd", reflect.TypeOf((*MockClient)(nil).CollaboratorAdd), arg0, arg1)
}

// CollaboratorRemove mocks base method
func (m *MockClient) CollaboratorRemove(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "CollaboratorRemove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CollaboratorRemove indicates an expected call of CollaboratorRemove
func (mr *MockClientMockRecorder) CollaboratorRemove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollaboratorRemove", reflect.TypeOf((*MockClient)(nil).CollaboratorRemove), arg0, arg1)
}

// CollaboratorsList mocks base method
func (m *MockClient) CollaboratorsList(arg0 string) ([]go_scalingo.Collaborator, error) {
	ret := m.ctrl.Call(m, "CollaboratorsList", arg0)
	ret0, _ := ret[0].([]go_scalingo.Collaborator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollaboratorsList indicates an expected call of CollaboratorsList
func (mr *MockClientMockRecorder) CollaboratorsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollaboratorsList", reflect.TypeOf((*MockClient)(nil).CollaboratorsList), arg0)
}

// CreateToken mocks base method
func (m *MockClient) CreateToken(arg0 go_scalingo.Token) (go_scalingo.Token, error) {
	ret := m.ctrl.Call(m, "CreateToken", arg0)
	ret0, _ := ret[0].(go_scalingo.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken
func (mr *MockClientMockRecorder) CreateToken(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockClient)(nil).CreateToken), arg0)
}

// Deployment mocks base method
func (m *MockClient) Deployment(arg0, arg1 string) (*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "Deployment", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deployment indicates an expected call of Deployment
func (mr *MockClientMockRecorder) Deployment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployment", reflect.TypeOf((*MockClient)(nil).Deployment), arg0, arg1)
}

// DeploymentList mocks base method
func (m *MockClient) DeploymentList(arg0 string) ([]*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "DeploymentList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentList indicates an expected call of DeploymentList
func (mr *MockClientMockRecorder) DeploymentList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentList", reflect.TypeOf((*MockClient)(nil).DeploymentList), arg0)
}

// DeploymentLogs mocks base method
func (m *MockClient) DeploymentLogs(arg0 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "DeploymentLogs", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentLogs indicates an expected call of DeploymentLogs
func (mr *MockClientMockRecorder) DeploymentLogs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentLogs", reflect.TypeOf((*MockClient)(nil).DeploymentLogs), arg0)
}

// DeploymentStream mocks base method
func (m *MockClient) DeploymentStream(arg0 string) (*websocket.Conn, error) {
	ret := m.ctrl.Call(m, "DeploymentStream", arg0)
	ret0, _ := ret[0].(*websocket.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentStream indicates an expected call of DeploymentStream
func (mr *MockClientMockRecorder) DeploymentStream(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentStream", reflect.TypeOf((*MockClient)(nil).DeploymentStream), arg0)
}

// DeploymentsCreate mocks base method
func (m *MockClient) DeploymentsCreate(arg0 string, arg1 *go_scalingo.DeploymentsCreateParams) (*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "DeploymentsCreate", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentsCreate indicates an expected call of DeploymentsCreate
func (mr *MockClientMockRecorder) DeploymentsCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentsCreate", reflect.TypeOf((*MockClient)(nil).DeploymentsCreate), arg0, arg1)
}

// DomainsAdd mocks base method
func (m *MockClient) DomainsAdd(arg0 string, arg1 go_scalingo.Domain) (go_scalingo.Domain, error) {
	ret := m.ctrl.Call(m, "DomainsAdd", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainsAdd indicates an expected call of DomainsAdd
func (mr *MockClientMockRecorder) DomainsAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainsAdd", reflect.TypeOf((*MockClient)(nil).DomainsAdd), arg0, arg1)
}

// DomainsList mocks base method
func (m *MockClient) DomainsList(arg0 string) ([]go_scalingo.Domain, error) {
	ret := m.ctrl.Call(m, "DomainsList", arg0)
	ret0, _ := ret[0].([]go_scalingo.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainsList indicates an expected call of DomainsList
func (mr *MockClientMockRecorder) DomainsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainsList", reflect.TypeOf((*MockClient)(nil).DomainsList), arg0)
}

// DomainsRemove mocks base method
func (m *MockClient) DomainsRemove(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "DomainsRemove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DomainsRemove indicates an expected call of DomainsRemove
func (mr *MockClientMockRecorder) DomainsRemove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainsRemove", reflect.TypeOf((*MockClient)(nil).DomainsRemove), arg0, arg1)
}

// DomainsUpdate mocks base method
func (m *MockClient) DomainsUpdate(arg0, arg1, arg2, arg3 string) (go_scalingo.Domain, error) {
	ret := m.ctrl.Call(m, "DomainsUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(go_scalingo.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainsUpdate indicates an expected call of DomainsUpdate
func (mr *MockClientMockRecorder) DomainsUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainsUpdate", reflect.TypeOf((*MockClient)(nil).DomainsUpdate), arg0, arg1, arg2, arg3)
}

// EventsList mocks base method
func (m *MockClient) EventsList(arg0 string, arg1 go_scalingo.PaginationOpts) (go_scalingo.Events, go_scalingo.PaginationMeta, error) {
	ret := m.ctrl.Call(m, "EventsList", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Events)
	ret1, _ := ret[1].(go_scalingo.PaginationMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EventsList indicates an expected call of EventsList
func (mr *MockClientMockRecorder) EventsList(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsList", reflect.TypeOf((*MockClient)(nil).EventsList), arg0, arg1)
}

// KeysAdd mocks base method
func (m *MockClient) KeysAdd(arg0, arg1 string) (*go_scalingo.Key, error) {
	ret := m.ctrl.Call(m, "KeysAdd", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeysAdd indicates an expected call of KeysAdd
func (mr *MockClientMockRecorder) KeysAdd(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysAdd", reflect.TypeOf((*MockClient)(nil).KeysAdd), arg0, arg1)
}

// KeysDelete mocks base method
func (m *MockClient) KeysDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "KeysDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// KeysDelete indicates an expected call of KeysDelete
func (mr *MockClientMockRecorder) KeysDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysDelete", reflect.TypeOf((*MockClient)(nil).KeysDelete), arg0)
}

// KeysList mocks base method
func (m *MockClient) KeysList() ([]go_scalingo.Key, error) {
	ret := m.ctrl.Call(m, "KeysList")
	ret0, _ := ret[0].([]go_scalingo.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeysList indicates an expected call of KeysList
func (mr *MockClientMockRecorder) KeysList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysList", reflect.TypeOf((*MockClient)(nil).KeysList))
}

// Login mocks base method
func (m *MockClient) Login(arg0, arg1 string) (*go_scalingo.LoginResponse, error) {
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockClientMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockClient)(nil).Login), arg0, arg1)
}

// Logs mocks base method
func (m *MockClient) Logs(arg0 string, arg1 int, arg2 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "Logs", arg0, arg1, arg2)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs
func (mr *MockClientMockRecorder) Logs(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockClient)(nil).Logs), arg0, arg1, arg2)
}

// LogsArchives mocks base method
func (m *MockClient) LogsArchives(arg0 string, arg1 int) (*go_scalingo.LogsArchivesResponse, error) {
	ret := m.ctrl.Call(m, "LogsArchives", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.LogsArchivesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogsArchives indicates an expected call of LogsArchives
func (mr *MockClientMockRecorder) LogsArchives(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsArchives", reflect.TypeOf((*MockClient)(nil).LogsArchives), arg0, arg1)
}

// LogsArchivesByCursor mocks base method
func (m *MockClient) LogsArchivesByCursor(arg0, arg1 string) (*go_scalingo.LogsArchivesResponse, error) {
	ret := m.ctrl.Call(m, "LogsArchivesByCursor", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.LogsArchivesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogsArchivesByCursor indicates an expected call of LogsArchivesByCursor
func (mr *MockClientMockRecorder) LogsArchivesByCursor(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsArchivesByCursor", reflect.TypeOf((*MockClient)(nil).LogsArchivesByCursor), arg0, arg1)
}

// LogsURL mocks base method
func (m *MockClient) LogsURL(arg0 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "LogsURL", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogsURL indicates an expected call of LogsURL
func (mr *MockClientMockRecorder) LogsURL(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsURL", reflect.TypeOf((*MockClient)(nil).LogsURL), arg0)
}

// NotificationDestroy mocks base method
func (m *MockClient) NotificationDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "NotificationDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotificationDestroy indicates an expected call of NotificationDestroy
func (mr *MockClientMockRecorder) NotificationDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationDestroy", reflect.TypeOf((*MockClient)(nil).NotificationDestroy), arg0, arg1)
}

// NotificationPlatformByName mocks base method
func (m *MockClient) NotificationPlatformByName(arg0 string) ([]*go_scalingo.NotificationPlatform, error) {
	ret := m.ctrl.Call(m, "NotificationPlatformByName", arg0)
	ret0, _ := ret[0].([]*go_scalingo.NotificationPlatform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationPlatformByName indicates an expected call of NotificationPlatformByName
func (mr *MockClientMockRecorder) NotificationPlatformByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationPlatformByName", reflect.TypeOf((*MockClient)(nil).NotificationPlatformByName), arg0)
}

// NotificationPlatformsList mocks base method
func (m *MockClient) NotificationPlatformsList() ([]*go_scalingo.NotificationPlatform, error) {
	ret := m.ctrl.Call(m, "NotificationPlatformsList")
	ret0, _ := ret[0].([]*go_scalingo.NotificationPlatform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationPlatformsList indicates an expected call of NotificationPlatformsList
func (mr *MockClientMockRecorder) NotificationPlatformsList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationPlatformsList", reflect.TypeOf((*MockClient)(nil).NotificationPlatformsList))
}

// NotificationProvision mocks base method
func (m *MockClient) NotificationProvision(arg0, arg1 string) (go_scalingo.NotificationRes, error) {
	ret := m.ctrl.Call(m, "NotificationProvision", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.NotificationRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationProvision indicates an expected call of NotificationProvision
func (mr *MockClientMockRecorder) NotificationProvision(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationProvision", reflect.TypeOf((*MockClient)(nil).NotificationProvision), arg0, arg1)
}

// NotificationUpdate mocks base method
func (m *MockClient) NotificationUpdate(arg0, arg1, arg2 string) (go_scalingo.NotificationRes, error) {
	ret := m.ctrl.Call(m, "NotificationUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(go_scalingo.NotificationRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationUpdate indicates an expected call of NotificationUpdate
func (mr *MockClientMockRecorder) NotificationUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationUpdate", reflect.TypeOf((*MockClient)(nil).NotificationUpdate), arg0, arg1, arg2)
}

// NotificationsList mocks base method
func (m *MockClient) NotificationsList(arg0 string) ([]*go_scalingo.Notification, error) {
	ret := m.ctrl.Call(m, "NotificationsList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotificationsList indicates an expected call of NotificationsList
func (mr *MockClientMockRecorder) NotificationsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotificationsList", reflect.TypeOf((*MockClient)(nil).NotificationsList), arg0)
}

// NotifierByID mocks base method
func (m *MockClient) NotifierByID(arg0, arg1 string) (*go_scalingo.Notifier, error) {
	ret := m.ctrl.Call(m, "NotifierByID", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifierByID indicates an expected call of NotifierByID
func (mr *MockClientMockRecorder) NotifierByID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifierByID", reflect.TypeOf((*MockClient)(nil).NotifierByID), arg0, arg1)
}

// NotifierDestroy mocks base method
func (m *MockClient) NotifierDestroy(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "NotifierDestroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifierDestroy indicates an expected call of NotifierDestroy
func (mr *MockClientMockRecorder) NotifierDestroy(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifierDestroy", reflect.TypeOf((*MockClient)(nil).NotifierDestroy), arg0, arg1)
}

// NotifierProvision mocks base method
func (m *MockClient) NotifierProvision(arg0, arg1 string, arg2 go_scalingo.NotifierParams) (*go_scalingo.Notifier, error) {
	ret := m.ctrl.Call(m, "NotifierProvision", arg0, arg1, arg2)
	ret0, _ := ret[0].(*go_scalingo.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifierProvision indicates an expected call of NotifierProvision
func (mr *MockClientMockRecorder) NotifierProvision(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifierProvision", reflect.TypeOf((*MockClient)(nil).NotifierProvision), arg0, arg1, arg2)
}

// NotifierUpdate mocks base method
func (m *MockClient) NotifierUpdate(arg0, arg1, arg2 string, arg3 go_scalingo.NotifierParams) (*go_scalingo.Notifier, error) {
	ret := m.ctrl.Call(m, "NotifierUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*go_scalingo.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifierUpdate indicates an expected call of NotifierUpdate
func (mr *MockClientMockRecorder) NotifierUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifierUpdate", reflect.TypeOf((*MockClient)(nil).NotifierUpdate), arg0, arg1, arg2, arg3)
}

// NotifiersList mocks base method
func (m *MockClient) NotifiersList(arg0 string) (go_scalingo.Notifiers, error) {
	ret := m.ctrl.Call(m, "NotifiersList", arg0)
	ret0, _ := ret[0].(go_scalingo.Notifiers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifiersList indicates an expected call of NotifiersList
func (mr *MockClientMockRecorder) NotifiersList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifiersList", reflect.TypeOf((*MockClient)(nil).NotifiersList), arg0)
}

// OperationsShow mocks base method
func (m *MockClient) OperationsShow(arg0, arg1 string) (*go_scalingo.Operation, error) {
	ret := m.ctrl.Call(m, "OperationsShow", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OperationsShow indicates an expected call of OperationsShow
func (mr *MockClientMockRecorder) OperationsShow(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperationsShow", reflect.TypeOf((*MockClient)(nil).OperationsShow), arg0, arg1)
}

// Run mocks base method
func (m *MockClient) Run(arg0 go_scalingo.RunOpts) (*go_scalingo.RunRes, error) {
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(*go_scalingo.RunRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockClientMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockClient)(nil).Run), arg0)
}

// Self mocks base method
func (m *MockClient) Self() (*go_scalingo.User, error) {
	ret := m.ctrl.Call(m, "Self")
	ret0, _ := ret[0].(*go_scalingo.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Self indicates an expected call of Self
func (mr *MockClientMockRecorder) Self() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Self", reflect.TypeOf((*MockClient)(nil).Self))
}

// ShowToken mocks base method
func (m *MockClient) ShowToken(arg0 int) (go_scalingo.Token, error) {
	ret := m.ctrl.Call(m, "ShowToken", arg0)
	ret0, _ := ret[0].(go_scalingo.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowToken indicates an expected call of ShowToken
func (mr *MockClientMockRecorder) ShowToken(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowToken", reflect.TypeOf((*MockClient)(nil).ShowToken), arg0)
}

// SignUp mocks base method
func (m *MockClient) SignUp(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SignUp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp
func (mr *MockClientMockRecorder) SignUp(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockClient)(nil).SignUp), arg0, arg1)
}

// SourcesCreate mocks base method
func (m *MockClient) SourcesCreate() (*go_scalingo.Source, error) {
	ret := m.ctrl.Call(m, "SourcesCreate")
	ret0, _ := ret[0].(*go_scalingo.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SourcesCreate indicates an expected call of SourcesCreate
func (mr *MockClientMockRecorder) SourcesCreate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourcesCreate", reflect.TypeOf((*MockClient)(nil).SourcesCreate))
}

// TokensList mocks base method
func (m *MockClient) TokensList() (go_scalingo.Tokens, error) {
	ret := m.ctrl.Call(m, "TokensList")
	ret0, _ := ret[0].(go_scalingo.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokensList indicates an expected call of TokensList
func (mr *MockClientMockRecorder) TokensList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokensList", reflect.TypeOf((*MockClient)(nil).TokensList))
}

// UpdateUser mocks base method
func (m *MockClient) UpdateUser(arg0 go_scalingo.UpdateUserParams) (*go_scalingo.User, error) {
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(*go_scalingo.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockClientMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockClient)(nil).UpdateUser), arg0)
}

// UserEventsList mocks base method
func (m *MockClient) UserEventsList(arg0 go_scalingo.PaginationOpts) (go_scalingo.Events, go_scalingo.PaginationMeta, error) {
	ret := m.ctrl.Call(m, "UserEventsList", arg0)
	ret0, _ := ret[0].(go_scalingo.Events)
	ret1, _ := ret[1].(go_scalingo.PaginationMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserEventsList indicates an expected call of UserEventsList
func (mr *MockClientMockRecorder) UserEventsList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserEventsList", reflect.TypeOf((*MockClient)(nil).UserEventsList), arg0)
}

// VariableMultipleSet mocks base method
func (m *MockClient) VariableMultipleSet(arg0 string, arg1 go_scalingo.Variables) (go_scalingo.Variables, int, error) {
	ret := m.ctrl.Call(m, "VariableMultipleSet", arg0, arg1)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VariableMultipleSet indicates an expected call of VariableMultipleSet
func (mr *MockClientMockRecorder) VariableMultipleSet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableMultipleSet", reflect.TypeOf((*MockClient)(nil).VariableMultipleSet), arg0, arg1)
}

// VariableSet mocks base method
func (m *MockClient) VariableSet(arg0, arg1, arg2 string) (*go_scalingo.Variable, int, error) {
	ret := m.ctrl.Call(m, "VariableSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*go_scalingo.Variable)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VariableSet indicates an expected call of VariableSet
func (mr *MockClientMockRecorder) VariableSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableSet", reflect.TypeOf((*MockClient)(nil).VariableSet), arg0, arg1, arg2)
}

// VariableUnset mocks base method
func (m *MockClient) VariableUnset(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "VariableUnset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VariableUnset indicates an expected call of VariableUnset
func (mr *MockClientMockRecorder) VariableUnset(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariableUnset", reflect.TypeOf((*MockClient)(nil).VariableUnset), arg0, arg1)
}

// VariablesList mocks base method
func (m *MockClient) VariablesList(arg0 string) (go_scalingo.Variables, error) {
	ret := m.ctrl.Call(m, "VariablesList", arg0)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VariablesList indicates an expected call of VariablesList
func (mr *MockClientMockRecorder) VariablesList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariablesList", reflect.TypeOf((*MockClient)(nil).VariablesList), arg0)
}

// VariablesListWithoutAlias mocks base method
func (m *MockClient) VariablesListWithoutAlias(arg0 string) (go_scalingo.Variables, error) {
	ret := m.ctrl.Call(m, "VariablesListWithoutAlias", arg0)
	ret0, _ := ret[0].(go_scalingo.Variables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VariablesListWithoutAlias indicates an expected call of VariablesListWithoutAlias
func (mr *MockClientMockRecorder) VariablesListWithoutAlias(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VariablesListWithoutAlias", reflect.TypeOf((*MockClient)(nil).VariablesListWithoutAlias), arg0)
}
