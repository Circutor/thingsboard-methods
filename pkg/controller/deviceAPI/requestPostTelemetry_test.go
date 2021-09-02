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

func TestFailBodyEncodeattrBody(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	deviceToken := "00000000000000000000"
	attrBody := map[string]interface{}{"typeCosPhi": 0.95, "cosPhiDaily" + "Type": "L"}

	data := new(dataMock.InterfaceDataMock)
	data.On("BodyEncode", attrBody).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := deviceapi.NewControllerDeviceAPIMock("/", "", "", data, nil)

	status, _, _ := controller.PostTelemetry(deviceToken, attrBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestPostTelemetry(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	deviceToken := "00000000000000000000"
	attrBody := map[string]interface{}{"typeCosPhi": 0.95, "cosPhiDaily" + "Type": "L"}

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(attrBody)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", attrBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/v1/" + deviceToken + "/telemetry"

	request.On("CreateNewRequest", "POST", URL, "", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := deviceapi.NewControllerDeviceAPIMock("/", "", "", data, request)

	status, _, _ := controller.PostTelemetry(deviceToken, attrBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}
