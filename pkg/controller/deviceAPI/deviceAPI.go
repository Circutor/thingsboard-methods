// Copyright (c) 2021 Circutor S.A. All rights reserved.

package deviceapi

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
	deviceAPI = "api/v1"
	// Methods.
	attributes = "/attributes"
	telemetry  = "/telemetry"
)

// ThingsBoardDeviceAPIController methods call API ThingsBoard.
type ThingsBoardDeviceAPIController interface {
	PostDeviceAttributes(deviceToken string, attrBody interface{}) (int, map[string]interface{}, error)
	GetDeviceAttributes(deviceToken string, query map[string]interface{}) (int, map[string]interface{}, error)
	PostTelemetry(deviceToken string, attrBody interface{}) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardDeviceAPIController --structname ThingsBoardDeviceAPIControllerMock --filename ThingsBoardDeviceAPIControllerMock.go

type ControllerDeviceAPI struct {
	TB      core.ThingsBoard
	Data    data.InterfaceData
	Request request.InterfaceRequest
}

// NewControllerDeviceAPI creates a new ThingsBoardAuthController.
func NewControllerDeviceAPI(urlServer, userName, userPassword string) ControllerDeviceAPI {
	dataLib := data.NewData()
	requestLib := request.NewRequest()

	return ControllerDeviceAPI{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    &dataLib,
		Request: &requestLib,
	}
}

//nolint:interfacer
// NewControllerDeviceAPIMock creates a new ThingsBoardAuthController.
func NewControllerDeviceAPIMock(urlServer, userName, userPassword string,
	data *dataMock.InterfaceDataMock, request *requestMock.InterfaceRequestMock) ControllerDeviceAPI {
	return ControllerDeviceAPI{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    data,
		Request: request,
	}
}
