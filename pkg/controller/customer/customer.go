// Copyright (c) 2021 Circutor S.A. All rights reserved.

package customer

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
	customer  = "api/customer"
	customers = "api/customers"
	// Methods.
	getTenantCustomer = "api/tenant/customers"
)

// ThingsBoardCustomerController methods call API ThingsBoard.
type ThingsBoardCustomerController interface {
	SaveCustomer(saveCustomerBody core.SaveCustomerBody, token string) (int, map[string]interface{}, error)
	GetCustomerByID(id, token string) (int, map[string]interface{}, error)
	GetTenantCustomer(email, token string) (int, map[string]interface{}, error)
	GetCustomers(token string, query map[string]interface{}) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardCustomerController --structname ThingsBoardCustomerControllerMock --filename ThingsBoardCustomerControllerMock.go

type ControllerCustomer struct {
	TB      core.ThingsBoard
	Data    data.InterfaceData
	Request request.InterfaceRequest
}

// NewControllerCustomer  creates a new ThingsBoardCustomerController.
func NewControllerCustomer(urlServer, userName, userPassword string) ControllerCustomer {
	dataLib := data.NewData()
	requestLib := request.NewRequest()

	return ControllerCustomer{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    &dataLib,
		Request: &requestLib,
	}
}

//nolint:interfacer
// NewControllerCustomerMock creates a new ThingsBoardCustomerController.
func NewControllerCustomerMock(urlServer, userName, userPassword string,
	data *dataMock.InterfaceDataMock, request *requestMock.InterfaceRequestMock) ControllerCustomer {
	return ControllerCustomer{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    data,
		Request: request,
	}
}
