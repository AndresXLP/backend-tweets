// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// Tweets is an autogenerated mock type for the Tweets type
type Tweets struct {
	mock.Mock
}

// CreateTweet provides a mock function with given fields: cntx
func (_m *Tweets) CreateTweet(cntx echo.Context) error {
	ret := _m.Called(cntx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(cntx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTweet provides a mock function with given fields: cntx
func (_m *Tweets) DeleteTweet(cntx echo.Context) error {
	ret := _m.Called(cntx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(cntx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTweets provides a mock function with given fields: cntx
func (_m *Tweets) GetTweets(cntx echo.Context) error {
	ret := _m.Called(cntx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(cntx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTweet provides a mock function with given fields: cntx
func (_m *Tweets) UpdateTweet(cntx echo.Context) error {
	ret := _m.Called(cntx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(cntx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTweets interface {
	mock.TestingT
	Cleanup(func())
}

// NewTweets creates a new instance of Tweets. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTweets(t mockConstructorTestingTNewTweets) *Tweets {
	mock := &Tweets{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
