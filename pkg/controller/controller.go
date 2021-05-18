// Copyright (c) 2021 Circutor S.A. All rights reserved.

package controller

import (
	"github.com/circutor/thingsboard-methods/pkg/controller/auth"
	AuthMock "github.com/circutor/thingsboard-methods/pkg/controller/auth/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/customer"
	CustomerMock "github.com/circutor/thingsboard-methods/pkg/controller/customer/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/device"
	deviceMock "github.com/circutor/thingsboard-methods/pkg/controller/device/mocks"
	deviceAPI "github.com/circutor/thingsboard-methods/pkg/controller/deviceAPI"
	deviceAPIMock "github.com/circutor/thingsboard-methods/pkg/controller/deviceAPI/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/signup"
	SignUpMock "github.com/circutor/thingsboard-methods/pkg/controller/signup/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/telemetry"
	TelemetryMock "github.com/circutor/thingsboard-methods/pkg/controller/telemetry/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/user"
	UserMock "github.com/circutor/thingsboard-methods/pkg/controller/user/mocks"
)

type ThingsBoardController struct {
	Auth      auth.ThingsBoardAuthController
	Device    device.ThingsBoardDeviceController
	DeviceAPI deviceAPI.ThingsBoardDeviceAPIController
	SignUp    signup.ThingsBoardSignUpController
	User      user.ThingsBoardUserController
	Customer  customer.ThingsBoardCustomerController
	Telemetry telemetry.ThingsBoardTelemetryController
}

// NewThingsBoardController creates a new ThingsBoardController.
func NewThingsBoardController(urlServer, userName, userPassword string) ThingsBoardController {
	authController := auth.NewControllerAuth(urlServer, userName, userPassword)
	deviceController := device.NewControllerDevice(urlServer, userName, userPassword)
	deviceAPIController := deviceAPI.NewControllerDeviceAPI(urlServer, userName, userPassword)
	signUpController := signup.NewControllerSignUp(urlServer, userName, userPassword)
	userController := user.NewControllerUser(urlServer, userName, userPassword)
	customerController := customer.NewControllerCustomer(urlServer, userName, userPassword)
	telemetryController := telemetry.NewControllerTelemetry(urlServer, userName, userPassword)

	controller := ThingsBoardController{
		Auth:      &authController,
		Device:    &deviceController,
		DeviceAPI: &deviceAPIController,
		SignUp:    &signUpController,
		User:      &userController,
		Customer:  &customerController,
		Telemetry: &telemetryController,
	}

	return controller
}

//nolint:interfacer
// NewThingsBoardControllerMock creates a new ThingsBoardController.
func NewThingsBoardControllerMock(
	auth *AuthMock.ThingsBoardAuthControllerMock, device *deviceMock.ThingsBoardDeviceControllerMock,
	deviceAPI *deviceAPIMock.ThingsBoardDeviceAPIControllerMock, signUp *SignUpMock.ThingsBoardSignUpControllerMock,
	user *UserMock.ThingsBoardUserControllerMock, customer *CustomerMock.ThingsBoardCustomerControllerMock,
	telemetry *TelemetryMock.ThingsBoardTelemetryControllerMock) ThingsBoardController {
	controller := ThingsBoardController{
		Auth:      auth,
		Device:    device,
		DeviceAPI: deviceAPI,
		SignUp:    signUp,
		User:      user,
		Customer:  customer,
		Telemetry: telemetry,
	}

	return controller
}
