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

// Resource provides a mock function with given fields: c
func (_m *Tweets) Resource(c *echo.Group) {
	_m.Called(c)
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
