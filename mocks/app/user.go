// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/andresxlp/backend-twitter/internal/domain/dto"
	mock "github.com/stretchr/testify/mock"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, newUser
func (_m *User) CreateUser(ctx context.Context, newUser dto.NewUser) error {
	ret := _m.Called(ctx, newUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.NewUser) error); ok {
		r0 = rf(ctx, newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUser(t mockConstructorTestingTNewUser) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}