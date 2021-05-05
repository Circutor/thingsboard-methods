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
)

// ThingsBoardTelemetryController methods call API ThingsBoard.
type ThingsBoardTelemetryController interface {
	SaveDeviceAttributes(deviceID, scope, token string,
		attributesBody map[string]interface{}) (int, map[string]interface{}, error)
	SaveEntityAttributesV2(entityType, entityID, scope, token string,
		attributesBody map[string]interface{}) (int, map[string]interface{}, error)
	GetAttributeKeys(entityType, entityID, token string) (int, map[string]interface{}, error)
	GetAttributeKeysByScope(entityType, entityID, scope, token string) (int, map[string]interface{}, error)
	GetAttributes(entityType, entityID, token string, query map[string]interface{}) (int, map[string]interface{}, error)
	GetAttributesByScope(entityType, entityID, scope, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
	DeleteEntityAttributes(entityType, entityID, scope, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
}

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