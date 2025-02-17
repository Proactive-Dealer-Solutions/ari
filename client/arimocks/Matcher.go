// Code generated by mockery v2.16.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/CyCoreSystems/ari/v6"
	mock "github.com/stretchr/testify/mock"
)

// Matcher is an autogenerated mock type for the Matcher type
type Matcher struct {
	mock.Mock
}

// Match provides a mock function with given fields: o
func (_m *Matcher) Match(o *ari.Key) bool {
	ret := _m.Called(o)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*ari.Key) bool); ok {
		r0 = rf(o)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewMatcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewMatcher creates a new instance of Matcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMatcher(t mockConstructorTestingTNewMatcher) *Matcher {
	mock := &Matcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
