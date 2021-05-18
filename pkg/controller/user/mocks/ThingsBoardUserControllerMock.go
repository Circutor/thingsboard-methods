// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	core "github.com/circutor/thingsboard-methods/pkg/core"
	mock "github.com/stretchr/testify/mock"
)

// ThingsBoardUserControllerMock is an autogenerated mock type for the ThingsBoardUserController type
type ThingsBoardUserControllerMock struct {
	mock.Mock
}

// SaveUser provides a mock function with given fields: saveUserBody, token
func (_m *ThingsBoardUserControllerMock) SaveUser(saveUserBody core.SaveUserBody, token string) (int, map[string]interface{}, error) {
	ret := _m.Called(saveUserBody, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.SaveUserBody, string) int); ok {
		r0 = rf(saveUserBody, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(core.SaveUserBody, string) map[string]interface{}); ok {
		r1 = rf(saveUserBody, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.SaveUserBody, string) error); ok {
		r2 = rf(saveUserBody, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
