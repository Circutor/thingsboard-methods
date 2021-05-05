// Copyright (c) 2021 Circutor S.A. All rights reserved.

package deviceapi

import "github.com/circutor/thingsboard-methods/pkg/core"

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	deviceAPI = "api/v1"
	// Methods.
	attributes = "/attributes"
)

// ThingsBoardDeviceAPIController methods call API ThingsBoard.
type ThingsBoardDeviceAPIController interface {
	PostDeviceAttributes(deviceToken string, attrBody interface{}) (int, map[string]interface{}, error)
	GetDeviceAttributes(deviceToken string, query map[string]interface{}) (int, map[string]interface{}, error)
}

type ControllerDeviceAPI struct {
	TB core.ThingsBoard
}

// NewControllerDeviceAPI creates a new ThingsBoardAuthController.
func NewControllerDeviceAPI(urlServer, userName, userPassword string) ControllerDeviceAPI {
	tb := ControllerDeviceAPI{
		TB: core.NewThingsBoard(urlServer, userName, userPassword),
	}

	return tb
}