// Copyright (c) 2021 Circutor S.A. All rights reserved.

package customer

import "github.com/circutor/thingsboard-methods/pkg/core"

const (
	thingsBoard = "ThingsBoard call: %s"
	// Controller.
	customer = "api/customer"
	// Methods.
	getTenantCustomer = "api/tenant/customers"
)

// ThingsBoardCustomerController methods call API ThingsBoard.
type ThingsBoardCustomerController interface {
	SaveCustomer(saveCustomerBody core.SaveCustomerBody, token string) (int, map[string]interface{}, error)
	GetCustomerByID(id, token string) (int, map[string]interface{}, error)
	GetTenantCustomer(email, token string) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardCustomerController --structname ThingsBoardCustomerControllerMock --filename ThingsBoardCustomerControllerMock.go

type ControllerCustomer struct {
	TB core.ThingsBoard
}

// NewControllerCustomer  creates a new ThingsBoardCustomerController.
func NewControllerCustomer(urlServer, userName, userPassword string) ControllerCustomer {
	tb := ControllerCustomer{
		TB: core.NewThingsBoard(urlServer, userName, userPassword),
	}

	return tb
}
