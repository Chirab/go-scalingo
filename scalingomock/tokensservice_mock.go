// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo/v4 (interfaces: TokensService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	context "context"
	reflect "reflect"

	scalingo "github.com/Scalingo/go-scalingo/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockTokensService is a mock of TokensService interface.
type MockTokensService struct {
	ctrl     *gomock.Controller
	recorder *MockTokensServiceMockRecorder
}

// MockTokensServiceMockRecorder is the mock recorder for MockTokensService.
type MockTokensServiceMockRecorder struct {
	mock *MockTokensService
}

// NewMockTokensService creates a new mock instance.
func NewMockTokensService(ctrl *gomock.Controller) *MockTokensService {
	mock := &MockTokensService{ctrl: ctrl}
	mock.recorder = &MockTokensServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokensService) EXPECT() *MockTokensServiceMockRecorder {
	return m.recorder
}

// TokenCreate mocks base method.
func (m *MockTokensService) TokenCreate(arg0 context.Context, arg1 scalingo.TokenCreateParams) (scalingo.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenCreate", arg0, arg1)
	ret0, _ := ret[0].(scalingo.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenCreate indicates an expected call of TokenCreate.
func (mr *MockTokensServiceMockRecorder) TokenCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenCreate", reflect.TypeOf((*MockTokensService)(nil).TokenCreate), arg0, arg1)
}

// TokenExchange mocks base method.
func (m *MockTokensService) TokenExchange(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenExchange", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenExchange indicates an expected call of TokenExchange.
func (mr *MockTokensServiceMockRecorder) TokenExchange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenExchange", reflect.TypeOf((*MockTokensService)(nil).TokenExchange), arg0, arg1)
}

// TokenShow mocks base method.
func (m *MockTokensService) TokenShow(arg0 context.Context, arg1 int) (scalingo.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenShow", arg0, arg1)
	ret0, _ := ret[0].(scalingo.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenShow indicates an expected call of TokenShow.
func (mr *MockTokensServiceMockRecorder) TokenShow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenShow", reflect.TypeOf((*MockTokensService)(nil).TokenShow), arg0, arg1)
}

// TokensList mocks base method.
func (m *MockTokensService) TokensList(arg0 context.Context) (scalingo.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokensList", arg0)
	ret0, _ := ret[0].(scalingo.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokensList indicates an expected call of TokensList.
func (mr *MockTokensServiceMockRecorder) TokensList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokensList", reflect.TypeOf((*MockTokensService)(nil).TokensList), arg0)
}
