// Copyright (c) 2021 Circutor S.A. All rights reserved.

package entitygroup

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
	groups = "api/entityGroups/"
)

// ThingsBoardEntityGroupController methods call API ThingsBoard.
type ThingsBoardEntityGroupController interface {
	GetEntityGroupsByType(groupType, token string) (int, []interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardEntityGroupController --structname ThingsBoardEntityGroupControllerMock --filename ThingsBoardEntityGroupControllerMock.go

type ControllerEntityGroup struct {
	TB      core.ThingsBoard
	Data    data.InterfaceData
	Request request.InterfaceRequest
}

// NewControllerEntityGroup creates a new ThingsBoardEntityGroupController.
func NewControllerEntityGroup(urlServer, userName, userPassword string) ControllerEntityGroup {
	dataLib := data.NewData()
	requestLib := request.NewRequest()

	return ControllerEntityGroup{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    &dataLib,
		Request: &requestLib,
	}
}

//nolint:interfacer
// NewControllerEntityGroupMock creates a new ThingsBoardEntityGroupController.
func NewControllerEntityGroupMock(urlServer, userName, userPassword string,
	data *dataMock.InterfaceDataMock, request *requestMock.InterfaceRequestMock) ControllerEntityGroup {
	return ControllerEntityGroup{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    data,
		Request: request,
	}
}
