// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth

import (
	"github.com/circutor/thingsboard-methods/pkg/controller/auth/mocks"
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

//go:generate mockery --name ThingsBoardAuthController --structname AuthControllerMock --filename AuthControllerMock.go

type ControllerAuth struct {
	TB core.ThingsBoard
}

// NewControllerAuth creates a new ThingsBoardAuthController.
func NewControllerAuth(urlServer, userName, userPassword string) ControllerAuth {
	tb := ControllerAuth{
		TB: core.NewThingsBoard(urlServer, userName, userPassword),
	}

	return tb
}

// NewControllerAuthMock creates a new AuthControllerMock.
func NewControllerAuthMock() *mocks.AuthControllerMock {
	mock := new(mocks.AuthControllerMock)

	return mock
}
