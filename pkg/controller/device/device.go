// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device

import "github.com/circutor/thingsboard-methods/pkg/core"

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	device = "api/device"
	// Methods.
	credentials = "/credentials"
	claimDevice = "api/customer/device/"
	claim       = "claim"
)

// ThingsBoardDeviceController methods call API ThingsBoard.
type ThingsBoardDeviceController interface {
	GetDeviceByID(deviceIds, token string) (int, map[string]interface{}, error)
	GetDeviceCredentialsByDeviceID(deviceID, token string) (int, map[string]interface{}, error)
	ClaimDevice(deviceName, token string, claimDeviceBody core.ClaimDeviceBody) (int, map[string]interface{}, error)
	CreateDevice(createDeviceBody core.CreateDeviceBody, token string) (int, map[string]interface{}, error)
}

type ControllerDevice struct {
	TB core.ThingsBoard
}

// NewControllerDevice creates a new ThingsBoardAuthController.
func NewControllerDevice(urlServer, userName, userPassword string) ControllerDevice {
	tb := ControllerDevice{
		TB: core.NewThingsBoard(urlServer, userName, userPassword),
	}

	return tb
}