// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import "github.com/circutor/thingsboard-methods/pkg/core"

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	telemetry = "api/plugins/telemetry/"
	// Methods.
	saveAttributes      = "/attributes/"
	getAttributesKeys   = "/keys/attributes"
	getAttributesValues = "/values/attributes"
	getTimeseriesValues = "/values/timeseries"
)

// ThingsBoardTelemetryController methods call API ThingsBoard.
type ThingsBoardTelemetryController interface {
	SaveDeviceAttributes(deviceID, scope, token string,
		attributesBody map[string]interface{}) (int, []interface{}, error)
	SaveEntityAttributesV2(entityType, entityID, scope, token string,
		attributesBody map[string]interface{}) (int, []interface{}, error)
	GetAttributeKeys(entityType, entityID, token string) (int, []interface{}, error)
	GetAttributeKeysByScope(entityType, entityID, scope, token string) (int, []interface{}, error)
	GetAttributes(entityType, entityID, token string, query map[string]interface{}) (int, []interface{}, error)
	GetAttributesByScope(entityType, entityID, scope, token string,
		query map[string]interface{}) (int, []interface{}, error)
	DeleteEntityAttributes(entityType, entityID, scope, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
	GetLatestTimeseries(entityType, entityID, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
	GetTimeseries(entityType, entityID, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardTelemetryController --structname ThingsBoardTelemetryControllerMock --filename ThingsBoardTelemetryControllerMock.go

type ControllerTelemetry struct {
	TB core.ThingsBoard
}

// NewControllerTelemetry creates a new ThingsBoardTelemetryController.
func NewControllerTelemetry(urlServer, userName, userPassword string) ControllerTelemetry {
	tb := ControllerTelemetry{
		TB: core.NewThingsBoard(urlServer, userName, userPassword),
	}

	return tb
}
