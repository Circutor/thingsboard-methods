// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device_test

import (
	"net/http"
	"testing"
	"time"

	dataCall "github.com/circutor/common-library/pkg/data"
	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/device"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailBodyEncodeCreateDevice(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	createDeviceBody := core.NewCreateDeviceBody(
		"00000000000000", "COMPUTERC6", time.Now().UnixNano()/int64(time.Millisecond))

	entityGroupID := map[string]interface{}{
		"entityGroupId": "00000000-0000-0000-0000-000000000000",
	}

	data := new(dataMock.InterfaceDataMock)
	data.On("BodyEncode", createDeviceBody).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := device.NewControllerDeviceMock("/", "", "", data, nil)

	status, _, _ := controller.CreateDevice(createDeviceBody, "Bearer token_value", entityGroupID)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestCreateDevice(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	createDeviceBody := core.NewCreateDeviceBody(
		"00000000000000", "COMPUTERC6", time.Now().UnixNano()/int64(time.Millisecond))
	entityGroupID := map[string]interface{}{
		"entityGroupId": "00000000-0000-0000-0000-000000000000",
	}

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(createDeviceBody)
	query := map[string]interface{}{
		"entityGroupId": "00000000-0000-0000-0000-000000000000",
	}

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", createDeviceBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/device"

	request.On("CreateNewRequest", "POST", URL, "Bearer token_value", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := device.NewControllerDeviceMock("/", "", "", data, request)

	status, _, _ := controller.CreateDevice(createDeviceBody, "Bearer token_value", entityGroupID)
	assert.Equal(t, http.StatusInternalServerError, status)
}
