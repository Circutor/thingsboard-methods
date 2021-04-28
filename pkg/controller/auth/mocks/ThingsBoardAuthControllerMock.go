// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	core "thingsboard-methods/pkg/core"

	mock "github.com/stretchr/testify/mock"
)

// ThingsBoardAuthControllerMock is an autogenerated mock type for the ThingsBoardAuthController type
type ThingsBoardAuthControllerMock struct {
	mock.Mock
}

// Login provides a mock function with given fields: loginBody
func (_m *ThingsBoardAuthControllerMock) Login(loginBody core.LoginBody) (int, map[string]interface{}, error) {
	ret := _m.Called(loginBody)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.LoginBody) int); ok {
		r0 = rf(loginBody)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(core.LoginBody) map[string]interface{}); ok {
		r1 = rf(loginBody)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.LoginBody) error); ok {
		r2 = rf(loginBody)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
