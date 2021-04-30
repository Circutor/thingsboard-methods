// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// ResetPasswordByEmailBody body for a ResetPasswordByEmail body.
type ResetPasswordByEmailBody struct {
	Email string `json:"email"`
}

// NewResetPasswordByEmailBody create a new ResetPasswordByEmailBody for a ResetPasswordByEmail.
func NewResetPasswordByEmailBody(email string) ResetPasswordByEmailBody {
	resetPasswordByEmailBody := ResetPasswordByEmailBody{
		Email: email,
	}

	return resetPasswordByEmailBody
}
