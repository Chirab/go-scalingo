// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo/v4 (interfaces: RunsService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	context "context"
	reflect "reflect"

	scalingo "github.com/Scalingo/go-scalingo/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockRunsService is a mock of RunsService interface.
type MockRunsService struct {
	ctrl     *gomock.Controller
	recorder *MockRunsServiceMockRecorder
}

// MockRunsServiceMockRecorder is the mock recorder for MockRunsService.
type MockRunsServiceMockRecorder struct {
	mock *MockRunsService
}

// NewMockRunsService creates a new mock instance.
func NewMockRunsService(ctrl *gomock.Controller) *MockRunsService {
	mock := &MockRunsService{ctrl: ctrl}
	mock.recorder = &MockRunsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunsService) EXPECT() *MockRunsServiceMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockRunsService) Run(arg0 context.Context, arg1 scalingo.RunOpts) (*scalingo.RunRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(*scalingo.RunRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockRunsServiceMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunsService)(nil).Run), arg0, arg1)
}
