// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	core "github.com/circutor/thingsboard-methods/pkg/core"

	mock "github.com/stretchr/testify/mock"
)

// ThingsBoardCustomerControllerMock is an autogenerated mock type for the ThingsBoardCustomerController type
type ThingsBoardCustomerControllerMock struct {
	mock.Mock
}

// GetCustomerByID provides a mock function with given fields: id, token
func (_m *ThingsBoardCustomerControllerMock) GetCustomerByID(id string, token string) (int, map[string]interface{}, error) {
	ret := _m.Called(id, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(id, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string) map[string]interface{}); ok {
		r1 = rf(id, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(id, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetCustomers provides a mock function with given fields: token, query
func (_m *ThingsBoardCustomerControllerMock) GetCustomers(token string, query map[string]interface{}) (int, map[string]interface{}, error) {
	ret := _m.Called(token, query)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) int); ok {
		r0 = rf(token, query)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(token, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, map[string]interface{}) error); ok {
		r2 = rf(token, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTenantCustomer provides a mock function with given fields: email, token
func (_m *ThingsBoardCustomerControllerMock) GetTenantCustomer(email string, token string) (int, map[string]interface{}, error) {
	ret := _m.Called(email, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(email, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string) map[string]interface{}); ok {
		r1 = rf(email, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveCustomer provides a mock function with given fields: saveCustomerBody, token
func (_m *ThingsBoardCustomerControllerMock) SaveCustomer(saveCustomerBody core.SaveCustomerBody, token string) (int, map[string]interface{}, error) {
	ret := _m.Called(saveCustomerBody, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.SaveCustomerBody, string) int); ok {
		r0 = rf(saveCustomerBody, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(core.SaveCustomerBody, string) map[string]interface{}); ok {
		r1 = rf(saveCustomerBody, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(core.SaveCustomerBody, string) error); ok {
		r2 = rf(saveCustomerBody, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
