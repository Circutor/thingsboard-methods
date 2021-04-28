// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth

import "thingsboard-methods/pkg/core"

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	auth = "api/auth"
	// Methods.
	login = "/login"
)

// ThingsBoardAuthController methods call API ThingsBoard.
type ThingsBoardAuthController interface {
	Login(loginBody core.LoginBody) (int, map[string]interface{}, error)
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
