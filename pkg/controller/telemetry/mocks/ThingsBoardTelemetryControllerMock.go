// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ThingsBoardTelemetryControllerMock is an autogenerated mock type for the ThingsBoardTelemetryController type
type ThingsBoardTelemetryControllerMock struct {
	mock.Mock
}

// DeleteEntityAttributes provides a mock function with given fields: entityType, entityID, scope, token, query
func (_m *ThingsBoardTelemetryControllerMock) DeleteEntityAttributes(entityType string, entityID string, scope string, token string, query map[string]interface{}) (int, map[string]interface{}, error) {
	ret := _m.Called(entityType, entityID, scope, token, query)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]interface{}) int); ok {
		r0 = rf(entityType, entityID, scope, token, query)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, string, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(entityType, entityID, scope, token, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, string, map[string]interface{}) error); ok {
		r2 = rf(entityType, entityID, scope, token, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAttributeKeys provides a mock function with given fields: entityType, entityID, token
func (_m *ThingsBoardTelemetryControllerMock) GetAttributeKeys(entityType string, entityID string, token string) (int, []interface{}, error) {
	ret := _m.Called(entityType, entityID, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string) int); ok {
		r0 = rf(entityType, entityID, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []interface{}
	if rf, ok := ret.Get(1).(func(string, string, string) []interface{}); ok {
		r1 = rf(entityType, entityID, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string) error); ok {
		r2 = rf(entityType, entityID, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAttributeKeysByScope provides a mock function with given fields: entityType, entityID, scope, token
func (_m *ThingsBoardTelemetryControllerMock) GetAttributeKeysByScope(entityType string, entityID string, scope string, token string) (int, map[string]interface{}, error) {
	ret := _m.Called(entityType, entityID, scope, token)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, string) int); ok {
		r0 = rf(entityType, entityID, scope, token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, string) map[string]interface{}); ok {
		r1 = rf(entityType, entityID, scope, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, string) error); ok {
		r2 = rf(entityType, entityID, scope, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAttributes provides a mock function with given fields: entityType, entityID, token, query
func (_m *ThingsBoardTelemetryControllerMock) GetAttributes(entityType string, entityID string, token string, query map[string]interface{}) (int, []interface{}, error) {
	ret := _m.Called(entityType, entityID, token, query)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]interface{}) int); ok {
		r0 = rf(entityType, entityID, token, query)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, map[string]interface{}) []interface{}); ok {
		r1 = rf(entityType, entityID, token, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, map[string]interface{}) error); ok {
		r2 = rf(entityType, entityID, token, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAttributesByScope provides a mock function with given fields: entityType, entityID, scope, token, query
func (_m *ThingsBoardTelemetryControllerMock) GetAttributesByScope(entityType string, entityID string, scope string, token string, query map[string]interface{}) (int, map[string]interface{}, error) {
	ret := _m.Called(entityType, entityID, scope, token, query)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]interface{}) int); ok {
		r0 = rf(entityType, entityID, scope, token, query)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, string, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(entityType, entityID, scope, token, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, string, map[string]interface{}) error); ok {
		r2 = rf(entityType, entityID, scope, token, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLatestTimeseries provides a mock function with given fields: entityType, entityID, token, query
func (_m *ThingsBoardTelemetryControllerMock) GetLatestTimeseries(entityType string, entityID string, token string, query map[string]interface{}) (int, map[string]interface{}, error) {
	ret := _m.Called(entityType, entityID, token, query)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]interface{}) int); ok {
		r0 = rf(entityType, entityID, token, query)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(entityType, entityID, token, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, map[string]interface{}) error); ok {
		r2 = rf(entityType, entityID, token, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTimeseries provides a mock function with given fields: entityType, entityID, token, query
func (_m *ThingsBoardTelemetryControllerMock) GetTimeseries(entityType string, entityID string, token string, query map[string]interface{}) (int, map[string]interface{}, error) {
	ret := _m.Called(entityType, entityID, token, query)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]interface{}) int); ok {
		r0 = rf(entityType, entityID, token, query)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(entityType, entityID, token, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, map[string]interface{}) error); ok {
		r2 = rf(entityType, entityID, token, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveDeviceAttributes provides a mock function with given fields: deviceID, scope, token, attributesBody
func (_m *ThingsBoardTelemetryControllerMock) SaveDeviceAttributes(deviceID string, scope string, token string, attributesBody map[string]interface{}) (int, []interface{}, error) {
	ret := _m.Called(deviceID, scope, token, attributesBody)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, map[string]interface{}) int); ok {
		r0 = rf(deviceID, scope, token, attributesBody)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 []interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, map[string]interface{}) []interface{}); ok {
		r1 = rf(deviceID, scope, token, attributesBody)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, map[string]interface{}) error); ok {
		r2 = rf(deviceID, scope, token, attributesBody)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveEntityAttributesV1 provides a mock function with given fields: entityType, entityID, scope, token, attributesBody
func (_m *ThingsBoardTelemetryControllerMock) SaveEntityAttributesV1(entityType string, entityID string, scope string, token string, attributesBody map[string]interface{}) (int, map[string]interface{}, error) {
	ret := _m.Called(entityType, entityID, scope, token, attributesBody)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]interface{}) int); ok {
		r0 = rf(entityType, entityID, scope, token, attributesBody)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, string, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(entityType, entityID, scope, token, attributesBody)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, string, map[string]interface{}) error); ok {
		r2 = rf(entityType, entityID, scope, token, attributesBody)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveEntityAttributesV2 provides a mock function with given fields: entityType, entityID, scope, token, attributesBody
func (_m *ThingsBoardTelemetryControllerMock) SaveEntityAttributesV2(entityType string, entityID string, scope string, token string, attributesBody map[string]interface{}) (int, map[string]interface{}, error) {
	ret := _m.Called(entityType, entityID, scope, token, attributesBody)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]interface{}) int); ok {
		r0 = rf(entityType, entityID, scope, token, attributesBody)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, string, map[string]interface{}) map[string]interface{}); ok {
		r1 = rf(entityType, entityID, scope, token, attributesBody)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, string, map[string]interface{}) error); ok {
		r2 = rf(entityType, entityID, scope, token, attributesBody)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
