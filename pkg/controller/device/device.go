// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device

import "github.com/circutor/thingsboard-methods/pkg/core"

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	device = "api/device"
	// Methods.
	credentials     = "/credentials"
	claimDevice     = "api/customer/device/"
	claim           = "claim"
	devicesCustomer = "api/customer/"
	devicesGroup    = "api/entityGroup/"
	devices         = "/devices"
	devicesName     = "api/tenant/devices"
)

// ThingsBoardDeviceController methods call API ThingsBoard.
type ThingsBoardDeviceController interface {
	GetDeviceByID(deviceIds, token string) (int, map[string]interface{}, error)
	GetDeviceCredentialsByDeviceID(deviceID, token string) (int, map[string]interface{}, error)
	ClaimDevice(deviceName, token string, claimDeviceBody core.ClaimDeviceBody) (int, map[string]interface{}, error)
	ReClaimDevice(deviceName, token string) (int, map[string]interface{}, error)
	CreateDevice(createDeviceBody core.CreateDeviceBody, token string) (int, map[string]interface{}, error)
	DeleteDevice(deviceID, token string) (int, map[string]interface{}, error)
	GetCustomerDevices(customerID, token string, query map[string]interface{}) (int, map[string]interface{}, error)
	GetTenantDevice(deviceName, token string) (int, map[string]interface{}, error)
	GetDevicesByEntityGroupID(entityGroupID, token string,
		query map[string]interface{}) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardDeviceController --structname ThingsBoardDeviceControllerMock --filename ThingsBoardDeviceControllerMock.go

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
