// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device_test

import (
	"net/http"
	"testing"

	dataCall "github.com/circutor/common-library/pkg/data"
	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/device"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailBodyEncodeClaimDevice(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	claimDeviceBody := core.NewClaimDeviceBody("secret_key")

	data := new(dataMock.InterfaceDataMock)
	data.On("BodyEncode", claimDeviceBody).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := device.NewControllerDeviceMock("/", "", "", data, nil)

	status, _, _ := controller.ClaimDevice("ComputerC", "Bearer token_value", claimDeviceBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestLogin(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	claimDeviceBody := core.NewClaimDeviceBody(
		"secret_key")
	deviceName := "COMPUTERC"

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(claimDeviceBody)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", claimDeviceBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/customer/device/" + deviceName + "/claim"

	request.On("CreateNewRequest", "POST", URL, "Bearer token_value", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := device.NewControllerDeviceMock("/", "", "", data, request)

	status, _, _ := controller.ClaimDevice(deviceName, "Bearer token_value", claimDeviceBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}
