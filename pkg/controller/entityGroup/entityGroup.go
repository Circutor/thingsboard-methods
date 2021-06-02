// Copyright (c) 2021 Circutor S.A. All rights reserved.

package entitygroup

import "github.com/circutor/thingsboard-methods/pkg/core"

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	groups = "api/entityGroups/"
)

// ThingsBoardEntityGroupController methods call API ThingsBoard.
type ThingsBoardEntityGroupController interface {
	GetEntityGroupsByType(groupType, token string) (int, []interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardEntityGroupController --structname ThingsBoardEntityGroupControllerMock --filename TThingsBoardEntityGroupControllerMock.go

type ControllerEntityGroup struct {
	TB core.ThingsBoard
}

// NewControllerEntityGroup creates a new ThingsBoardEntityGroupController.
func NewControllerEntityGroup(urlServer, userName, userPassword string) ControllerEntityGroup {
	tb := ControllerEntityGroup{
		TB: core.NewThingsBoard(urlServer, userName, userPassword),
	}

	return tb
}
