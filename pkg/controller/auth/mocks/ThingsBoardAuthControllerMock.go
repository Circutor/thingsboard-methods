// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	core "github.com/circutor/thingsboard-methods/pkg/core"
	mock "github.com/stretchr/testify/mock"
)

// ThingsBoardAuthControllerMock is an autogenerated mock type for the ThingsBoardAuthController type
type ThingsBoardAuthControllerMock struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: changePasswordBody, token
func (_m *ThingsBoardAuthControllerMock) ChangePassword(changePasswordBody core.ChangePasswordBody, token string) (int, map[string]interface{}, error) {
	ret := _m.Called(changePasswordBody, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.ChangePasswordBody, string) int); ok {
		r0 = rf(changePasswordBody, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(core.ChangePasswordBody, string) map[string]interface{}); ok {
		r1 = rf(changePasswordBody, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.ChangePasswordBody, string) error); ok {
		r2 = rf(changePasswordBody, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUser provides a mock function with given fields: token
func (_m *ThingsBoardAuthControllerMock) GetUser(token string) (int, map[string]interface{}, error) {
	ret := _m.Called(token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string) map[string]interface{}); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
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

// Logout provides a mock function with given fields: token
func (_m *ThingsBoardAuthControllerMock) Logout(token string) (int, map[string]interface{}, error) {
	ret := _m.Called(token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string) map[string]interface{}); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RefreshToken provides a mock function with given fields: refreshTokenBody
func (_m *ThingsBoardAuthControllerMock) RefreshToken(refreshTokenBody core.RefreshTokenBody) (int, map[string]interface{}, error) {
	ret := _m.Called(refreshTokenBody)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.RefreshTokenBody) int); ok {
		r0 = rf(refreshTokenBody)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(core.RefreshTokenBody) map[string]interface{}); ok {
		r1 = rf(refreshTokenBody)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.RefreshTokenBody) error); ok {
		r2 = rf(refreshTokenBody)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
