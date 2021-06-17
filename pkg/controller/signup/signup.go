// Copyright (c) 2021 Circutor S.A. All rights reserved.

package signup

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
	noauth = "api/noauth"
	// Methods.
	signUp               = "/signup"
	activateUser         = "/activateByEmailCode"
	resetPasswordByEmail = "/resetPasswordByEmail"
	resetPassword        = "/resetPassword"
)

// ThingsBoardSignUpController methods call API ThingsBoard.
type ThingsBoardSignUpController interface {
	SignUp(signUpBody core.SignUpBody) (int, map[string]interface{}, error)
	ActivateUserByEmailCode(emailCode string) (int, map[string]interface{}, error)
	ResetPasswordByEmail(resetPasswordByEmailBody core.ResetPasswordByEmailBody,
		token string) (int, map[string]interface{}, error)
	ResetPassword(resetPasswordBody core.ResetPasswordBody) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name ThingsBoardSignUpController --structname ThingsBoardSignUpControllerMock --filename ThingsBoardSignUpControllerMock.go

type ControllerSignUp struct {
	TB      core.ThingsBoard
	Data    data.InterfaceData
	Request request.InterfaceRequest
}

// NewControllerSignUp creates a new ThingsBoardSignUpController.
func NewControllerSignUp(urlServer, userName, userPassword string) ControllerSignUp {
	dataLib := data.NewData()
	requestLib := request.NewRequest()

	return ControllerSignUp{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    &dataLib,
		Request: &requestLib,
	}
}

//nolint:interfacer
// NewControllerSignUpMock creates a new ThingsBoardSignUpController.
func NewControllerSignUpMock(urlServer, userName, userPassword string,
	data *dataMock.InterfaceDataMock, request *requestMock.InterfaceRequestMock) ControllerSignUp {
	return ControllerSignUp{
		TB:      core.NewThingsBoard(urlServer, userName, userPassword),
		Data:    data,
		Request: request,
	}
}
