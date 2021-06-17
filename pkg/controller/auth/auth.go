// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth

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
	auth = "api/auth"
	// Methods.
	login          = "/login"
	logout         = "/logout"
	refreshToken   = "/token"
	getUser        = "/user"
	changePassword = "/changePassword"
)

// ThingsBoardAuthController methods call API ThingsBoard.
type ThingsBoardAuthController interface {
	Login(loginBody core.LoginBody) (int, map[string]interface{}, error)
	Logout(token string) (int, map[string]interface{}, error)
	RefreshToken(refreshTokenBody core.RefreshTokenBody) (int, map[string]interface{}, error)
	GetUser(token string) (int, map[string]interface{}, error)
	ChangePassword(changePasswordBody core.ChangePasswordBody, token string) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardAuthController --structname ThingsBoardAuthControllerMock --filename ThingsBoardAuthControllerMock.go

type ControllerAuth struct {
	TB      core.ThingsBoard
	Data    data.InterfaceData
	Request request.InterfaceRequest
}

// NewControllerAuth creates a new ThingsBoardAuthController.
func NewControllerAuth(urlServer, userName, userPassword string) ControllerAuth {
	dataLib := data.NewData()
	requestLib := request.NewRequest()

	return ControllerAuth{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    &dataLib,
		Request: &requestLib,
	}
}

//nolint:interfacer
// NewControllerAuthMock creates a new ThingsBoardAuthController.
func NewControllerAuthMock(urlServer, userName, userPassword string,
	data *dataMock.InterfaceDataMock, request *requestMock.InterfaceRequestMock) ControllerAuth {
	return ControllerAuth{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    data,
		Request: request,
	}
}
