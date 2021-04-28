// Copyright (c) 2021 Circutor S.A. All rights reserved.

package controller

import "github.com/circutor/thingsboard-methods/pkg/controller/auth"

type ThingsBoardController struct {
	Auth auth.ControllerAuth
}

// NewThingsBoardController creates a new NewThingsBoardController.
func NewThingsBoardController(urlServer, userName, userPassword string) ThingsBoardController {
	controller := ThingsBoardController{
		Auth: auth.NewControllerAuth(urlServer, userName, userPassword),
	}

	return controller
}
