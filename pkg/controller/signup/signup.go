// Copyright (c) 2021 Circutor S.A. All rights reserved.

package signup

import "github.com/circutor/thingsboard-methods/pkg/core"

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

type ControllerSignUp struct {
	TB core.ThingsBoard
}

// NewControllerSignUp creates a new ThingsBoardSignUpController.
func NewControllerSignUp(urlServer, userName, userPassword string) ControllerSignUp {
	tb := ControllerSignUp{
		TB: core.NewThingsBoard(urlServer, userName, userPassword),
	}

	return tb
}
