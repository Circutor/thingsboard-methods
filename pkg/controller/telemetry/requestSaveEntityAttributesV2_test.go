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

func TestFailBodyEncodeSaveEntityAttributesV2(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	// Send Telemetry mock
	entityType := "USER"
	entityID := "00000000-0000-0000-0000-000000000000"
	scope := "SERVER_SCOPE"
	attrBody := map[string]interface{}{
		"attribute1": "A1",
	}

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", attrBody).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := telemetry.NewControllerTelemetryMock("/", "", "", data, nil)

	status, _, _ := controller.SaveEntityAttributesV2(
		entityType, entityID, scope, "Bearer token_value", attrBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestSaveEntityAttributesV2(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	// Send Telemetry mock
	entityType := "USER"
	entityID := "00000000-0000-0000-0000-000000000000"
	scope := "SERVER_SCOPE"
	attrBody := map[string]interface{}{
		"attribute1": "A1",
	}

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(attrBody)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", attrBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	request.On("CreateNewRequest", "POST",
		"/api/plugins/telemetry/"+entityType+"/"+entityID+"/attributes/"+scope, "Bearer token_value", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := telemetry.NewControllerTelemetryMock("/", "", "", data, request)

	status, _, _ := controller.SaveEntityAttributesV2(entityType, entityID, scope, "Bearer token_value", attrBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}
