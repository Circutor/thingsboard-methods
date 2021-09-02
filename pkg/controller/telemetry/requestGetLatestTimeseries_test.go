// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry_test

import (
	"net/http"
	"testing"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/telemetry"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailRequestGetLatestTimeseries(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	// Send Telemetry mock
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/plugins/telemetry/DEVICE/" + "00000000-0000-0000-0000-000000000000" + "/values/timeseries"

	request.On("CreateNewRequest", "GET", URL, "Bearer token_value", nil, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := telemetry.NewControllerTelemetryMock("/", "", "", data, request)

	status, _, _ := controller.GetLatestTimeseries(
		"DEVICE", "00000000-0000-0000-0000-000000000000", "Bearer token_value", query)
	assert.Equal(t, http.StatusInternalServerError, status)
}
