// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/odanraujo/financial-organizer-api/internal/repository/users (interfaces: UserRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	excp "github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	users "github.com/odanraujo/financial-organizer-api/internal/entity/users"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(arg0 context.Context, arg1 users.CreateUser) (users.CreateUser, *excp.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(users.CreateUser)
	ret1, _ := ret[1].(*excp.Exception)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), arg0, arg1)
}

// GetUserForCPF mocks base method.
func (m *MockUserRepository) GetUserForCPF(arg0 context.Context, arg1 string) (users.CreateUser, *excp.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserForCPFOrEmail", arg0, arg1)
	ret0, _ := ret[0].(users.CreateUser)
	ret1, _ := ret[1].(*excp.Exception)
	return ret0, ret1
}

// GetUserForCPF indicates an expected call of GetUserForCPF.
func (mr *MockUserRepositoryMockRecorder) GetUserForCPF(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserForCPFOrEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserForCPF), arg0, arg1)
}

// GetUserForEmail mocks base method.
func (m *MockUserRepository) GetUserForEmail(arg0 context.Context, arg1 string) (users.CreateUser, *excp.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserForEmail", arg0, arg1)
	ret0, _ := ret[0].(users.CreateUser)
	ret1, _ := ret[1].(*excp.Exception)
	return ret0, ret1
}

// GetUserForEmail indicates an expected call of GetUserForEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserForEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserForEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserForEmail), arg0, arg1)
}

// UpdateUserForCPF mocks base method.
func (m *MockUserRepository) UpdateUserForCPF(arg0 context.Context, arg1 string, arg2 users.CreateUser) (users.CreateUser, *excp.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserForCPF", arg0, arg1, arg2)
	ret0, _ := ret[0].(users.CreateUser)
	ret1, _ := ret[1].(*excp.Exception)
	return ret0, ret1
}

// UpdateUserForCPF indicates an expected call of UpdateUserForCPF.
func (mr *MockUserRepositoryMockRecorder) UpdateUserForCPF(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserForCPF", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserForCPF), arg0, arg1, arg2)
}

// UpdateUserForEmail mocks base method.
func (m *MockUserRepository) UpdateUserForEmail(arg0 context.Context, arg1 string, arg2 users.CreateUser) (users.CreateUser, *excp.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserForEmail", arg0, arg1, arg2)
	ret0, _ := ret[0].(users.CreateUser)
	ret1, _ := ret[1].(*excp.Exception)
	return ret0, ret1
}

// UpdateUserForEmail indicates an expected call of UpdateUserForEmail.
func (mr *MockUserRepositoryMockRecorder) UpdateUserForEmail(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserForEmail", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserForEmail), arg0, arg1, arg2)
}
