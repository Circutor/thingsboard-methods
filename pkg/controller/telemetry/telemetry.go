// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"github.com/circutor/common-library/pkg/data"
	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/request"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/core"
)

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	telemetry = "api/plugins/telemetry/"
	// Methods.
	saveAttributes      = "/attributes/"
	getAttributesKeys   = "/keys/attributes"
	getAttributesValues = "/values/attributes"
	getTimeseriesValues = "/values/timeseries"
	getTimeseriesDelete = "/timeseries/delete"
)

// ThingsBoardTelemetryController methods call API ThingsBoard.
type ThingsBoardTelemetryController interface {
	SaveDeviceAttributes(deviceID, scope, token string,
		attributesBody map[string]interface{}) (int, []interface{}, error)
	SaveEntityAttributesV1(entityType, entityID, scope, token string,
		attributesBody map[string]interface{}) (int, map[string]interface{}, error)
	SaveEntityAttributesV2(entityType, entityID, scope, token string,
		attributesBody map[string]interface{}) (int, map[string]interface{}, error)
	GetAttributeKeys(entityType, entityID, token string) (int, []interface{}, error)
	GetAttributeKeysByScope(entityType, entityID, scope, token string) (int, map[string]interface{}, error)
	GetAttributes(entityType, entityID, token string, query map[string]interface{}) (int, map[string]interface{}, error)
	GetAttributesByScope(entityType, entityID, scope, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
	DeleteEntityAttributes(entityType, entityID, scope, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
	GetLatestTimeseries(entityType, entityID, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
	GetTimeseries(entityType, entityID, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
	DeleteEntityTimeseries(entityType, entityID, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardTelemetryController --structname ThingsBoardTelemetryControllerMock --filename ThingsBoardTelemetryControllerMock.go

type ControllerTelemetry struct {
	TB      core.ThingsBoard
	Data    data.InterfaceData
	Request request.InterfaceRequest
}

// NewControllerTelemetry creates a new ThingsBoardTelemetryController.
func NewControllerTelemetry(urlServer, userName, userPassword string) ControllerTelemetry {
	dataLib := data.NewData()
	requestLib := request.NewRequest()

	return ControllerTelemetry{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    &dataLib,
		Request: &requestLib,
	}
}

//nolint:interfacer
// NewControllerTelemetryMock creates a new ThingsBoardTelemetryController.
func NewControllerTelemetryMock(urlServer, userName, userPassword string,
	data *dataMock.InterfaceDataMock, request *requestMock.InterfaceRequestMock) ControllerTelemetry {
	return ControllerTelemetry{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    data,
		Request: request,
	}
}
