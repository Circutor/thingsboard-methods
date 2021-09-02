// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry_test

import (
	"net/http"
	"testing"

	dataCall "github.com/circutor/common-library/pkg/data"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/telemetry"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailRequestDeleteEntityTimeseries(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	// Send Telemetry mock
	attrBody := map[string]interface{}{"keys": "targetCosPhi"}

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(attrBody)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", attrBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	request.On("CreateNewRequest", "DELETE",
		"/api/plugins/telemetry/USER/"+"00000000-0000-0000-0000-000000000000"+"/timeseries/delete",
		"Bearer token_value", nil, attrBody).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := telemetry.NewControllerTelemetryMock("/", "", "", data, request)

	status, _, _ := controller.DeleteEntityTimeseries("USER",
		"00000000-0000-0000-0000-000000000000", "Bearer token_value", attrBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}
