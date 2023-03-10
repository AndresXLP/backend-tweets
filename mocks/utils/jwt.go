// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// JWT is an autogenerated mock type for the JWT type
type JWT struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: data
func (_m *JWT) GenerateToken(data string) (string, error) {
	ret := _m.Called(data)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: receivedToken
func (_m *JWT) ValidateToken(receivedToken string) (string, error) {
	ret := _m.Called(receivedToken)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(receivedToken)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(receivedToken)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(receivedToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewJWT interface {
	mock.TestingT
	Cleanup(func())
}

// NewJWT creates a new instance of JWT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJWT(t mockConstructorTestingTNewJWT) *JWT {
	mock := &JWT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
