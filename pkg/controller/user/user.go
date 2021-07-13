// Copyright (c) 2021 Circutor S.A. All rights reserved.

package user

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
	user = "api/user"
)

// ThingsBoardUserController methods call API ThingsBoard.
type ThingsBoardUserController interface {
	SaveUser(saveUserBody core.SaveUserBody, token string) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardUserController --structname ThingsBoardUserControllerMock --filename ThingsBoardUserControllerMock.go

type ControllerUser struct {
	TB      core.ThingsBoard
	Data    data.InterfaceData
	Request request.InterfaceRequest
}

// NewControllerUser creates a new ThingsBoardUserController.
func NewControllerUser(urlServer, userName, userPassword string) ControllerUser {
	dataLib := data.NewData()
	requestLib := request.NewRequest()

	return ControllerUser{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    &dataLib,
		Request: &requestLib,
	}
}

//nolint:interfacer
// NewControllerUserMock creates a new ThingsBoardUserController.
func NewControllerUserMock(urlServer, userName, userPassword string,
	data *dataMock.InterfaceDataMock, request *requestMock.InterfaceRequestMock) ControllerUser {
	return ControllerUser{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    data,
		Request: request,
	}
}
