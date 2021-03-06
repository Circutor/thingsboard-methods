// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	core "github.com/circutor/thingsboard-methods/pkg/core"
	mock "github.com/stretchr/testify/mock"
)

// ThingsBoardSignUpControllerMock is an autogenerated mock type for the ThingsBoardSignUpController type
type ThingsBoardSignUpControllerMock struct {
	mock.Mock
}

// ActivateUserByEmailCode provides a mock function with given fields: emailCode
func (_m *ThingsBoardSignUpControllerMock) ActivateUserByEmailCode(emailCode string) (int, map[string]interface{}, error) {
	ret := _m.Called(emailCode)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(emailCode)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string) map[string]interface{}); ok {
		r1 = rf(emailCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(emailCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResetPassword provides a mock function with given fields: resetPasswordBody
func (_m *ThingsBoardSignUpControllerMock) ResetPassword(resetPasswordBody core.ResetPasswordBody) (int, map[string]interface{}, error) {
	ret := _m.Called(resetPasswordBody)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.ResetPasswordBody) int); ok {
		r0 = rf(resetPasswordBody)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(core.ResetPasswordBody) map[string]interface{}); ok {
		r1 = rf(resetPasswordBody)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.ResetPasswordBody) error); ok {
		r2 = rf(resetPasswordBody)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResetPasswordByEmail provides a mock function with given fields: resetPasswordByEmailBody, token
func (_m *ThingsBoardSignUpControllerMock) ResetPasswordByEmail(resetPasswordByEmailBody core.ResetPasswordByEmailBody, token string) (int, map[string]interface{}, error) {
	ret := _m.Called(resetPasswordByEmailBody, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.ResetPasswordByEmailBody, string) int); ok {
		r0 = rf(resetPasswordByEmailBody, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(core.ResetPasswordByEmailBody, string) map[string]interface{}); ok {
		r1 = rf(resetPasswordByEmailBody, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.ResetPasswordByEmailBody, string) error); ok {
		r2 = rf(resetPasswordByEmailBody, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SignUp provides a mock function with given fields: signUpBody
func (_m *ThingsBoardSignUpControllerMock) SignUp(signUpBody core.SignUpBody) (int, map[string]interface{}, error) {
	ret := _m.Called(signUpBody)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.SignUpBody) int); ok {
		r0 = rf(signUpBody)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(core.SignUpBody) map[string]interface{}); ok {
		r1 = rf(signUpBody)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.SignUpBody) error); ok {
		r2 = rf(signUpBody)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
