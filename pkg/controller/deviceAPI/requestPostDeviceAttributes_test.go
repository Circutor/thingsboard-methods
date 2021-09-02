// Copyright (c) 2021 Circutor S.A. All rights reserved.

package deviceapi_test

import (
	"net/http"
	"testing"

	dataCall "github.com/circutor/common-library/pkg/data"
	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	deviceapi "github.com/circutor/thingsboard-methods/pkg/controller/deviceAPI"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailBodyEncodePostDeviceAttributes(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	deviceToken := "00000000000000000000"
	bodyVertical := map[string]interface{}{"targetCosPhi": 0.95, "power": 0, "voltage": 400, "rejectionFilters": false}

	data := new(dataMock.InterfaceDataMock)
	data.On("BodyEncode", bodyVertical).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := deviceapi.NewControllerDeviceAPIMock("/", "", "", data, nil)

	status, _, _ := controller.PostDeviceAttributes(deviceToken, bodyVertical)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestPostDeviceAttributes(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	deviceToken := "00000000000000000000"
	bodyVertical := map[string]interface{}{"targetCosPhi": 0.95, "power": 0, "voltage": 400, "rejectionFilters": false}

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(bodyVertical)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", bodyVertical).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/v1/" + deviceToken + "/attributes"

	request.On("CreateNewRequest", "POST", URL, "", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := deviceapi.NewControllerDeviceAPIMock("/", "", "", data, request)

	status, _, _ := controller.PostDeviceAttributes(deviceToken, bodyVertical)
	assert.Equal(t, http.StatusInternalServerError, status)
}
