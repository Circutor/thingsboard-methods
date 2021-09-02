// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device_test

import (
	"net/http"
	"testing"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/device"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailRequestReClaimDevice(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	deviceName := "COMPUTERC6"

	query := map[string]interface{}(nil)
	data := new(dataMock.InterfaceDataMock)

	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/customer/device/" + deviceName + "/claim"

	request.On("CreateNewRequest", "DELETE", URL, "Bearer token_value", nil, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := device.NewControllerDeviceMock("/", "", "", data, request)

	status, _, _ := controller.ReClaimDevice(deviceName, "Bearer token_value")
	assert.Equal(t, http.StatusInternalServerError, status)
}
