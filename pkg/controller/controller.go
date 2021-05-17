// Copyright (c) 2021 Circutor S.A. All rights reserved.

package controller

import (
	"github.com/circutor/thingsboard-methods/pkg/controller/auth"
	"github.com/circutor/thingsboard-methods/pkg/controller/auth/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/customer"
	"github.com/circutor/thingsboard-methods/pkg/controller/device"
	deviceAPI "github.com/circutor/thingsboard-methods/pkg/controller/deviceAPI"
	"github.com/circutor/thingsboard-methods/pkg/controller/signup"
	"github.com/circutor/thingsboard-methods/pkg/controller/telemetry"
	"github.com/circutor/thingsboard-methods/pkg/controller/user"
)

type ThingsBoardController struct {
	Auth      auth.ControllerAuth
	Device    device.ControllerDevice
	DeviceAPI deviceAPI.ControllerDeviceAPI
	SignUp    signup.ControllerSignUp
	User      user.ControllerUser
	Customer  customer.ControllerCustomer
	Telemetry telemetry.ControllerTelemetry
}

type ThingsBoardControllerMock struct {
	Auth *mocks.AuthControllerMock
}

// NewThingsBoardController creates a new NewThingsBoardController.
func NewThingsBoardController(urlServer, userName, userPassword string) ThingsBoardController {
	controller := ThingsBoardController{
		Auth:      auth.NewControllerAuth(urlServer, userName, userPassword),
		Device:    device.NewControllerDevice(urlServer, userName, userPassword),
		DeviceAPI: deviceAPI.NewControllerDeviceAPI(urlServer, userName, userPassword),
		SignUp:    signup.NewControllerSignUp(urlServer, userName, userPassword),
		User:      user.NewControllerUser(urlServer, userName, userPassword),
		Customer:  customer.NewControllerCustomer(urlServer, userName, userPassword),
		Telemetry: telemetry.NewControllerTelemetry(urlServer, userName, userPassword),
	}

	return controller
}

// NewThingsBoardControllerMock create a new ThingsBoardControllerMock.
func NewThingsBoardControllerMock() ThingsBoardControllerMock {
	controller := ThingsBoardControllerMock{
		Auth: auth.NewControllerAuthMock(),
	}

	return controller
}
