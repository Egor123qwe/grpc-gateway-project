// Code generated by MockGen. DO NOT EDIT.
// Source: scenarios.go

// Package mock_scenarios is a generated GoMock package.
package mock_scenarios

import (
	context "context"
	reflect "reflect"

	models "github.com/Egor123qwe/grpc-gateway-project/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(ctx context.Context, usr *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, usr)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(ctx, usr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), ctx, usr)
}

// DeleteUser mocks base method.
func (m *MockUser) DeleteUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUser)(nil).DeleteUser), ctx, id)
}

// GetUser mocks base method.
func (m *MockUser) GetUser(ctx context.Context, id string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUser)(nil).GetUser), ctx, id)
}

// GetUserByToken mocks base method.
func (m *MockUser) GetUserByToken(ctx context.Context, token string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByToken", ctx, token)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByToken indicates an expected call of GetUserByToken.
func (mr *MockUserMockRecorder) GetUserByToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByToken", reflect.TypeOf((*MockUser)(nil).GetUserByToken), ctx, token)
}

// SubscribeUser mocks base method.
func (m *MockUser) SubscribeUser(ctx context.Context, ids *models.SubscribeEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeUser", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeUser indicates an expected call of SubscribeUser.
func (mr *MockUserMockRecorder) SubscribeUser(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeUser", reflect.TypeOf((*MockUser)(nil).SubscribeUser), ctx, ids)
}

// UnsubscribeUser mocks base method.
func (m *MockUser) UnsubscribeUser(ctx context.Context, ids *models.SubscribeEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeUser", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeUser indicates an expected call of UnsubscribeUser.
func (mr *MockUserMockRecorder) UnsubscribeUser(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeUser", reflect.TypeOf((*MockUser)(nil).UnsubscribeUser), ctx, ids)
}
