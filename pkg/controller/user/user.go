// Copyright (c) 2021 Circutor S.A. All rights reserved.

package user

import "github.com/circutor/thingsboard-methods/pkg/core"

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	user = "api/user"
)

// ThingsBoardUserController methods call API ThingsBoard.
type ThingsBoardUserController interface {
	SaveUser(saveUserBody core.SaveUserBody, token string) (int, map[string]interface{}, error)
}

type ControllerUser struct {
	TB core.ThingsBoard
}

// NewControllerSignUp creates a new ThingsBoardUserController.
func NewControllerUser(urlServer, userName, userPassword string) ControllerUser {
	tb := ControllerUser{
		TB: core.NewThingsBoard(urlServer, userName, userPassword),
	}

	return tb
}
