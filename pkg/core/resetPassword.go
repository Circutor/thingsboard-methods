// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// ResetPasswordBody  body for request resetPassword.
type ResetPasswordBody struct {
	ResetToken string `json:"resetToken"`
	Password   string `json:"password"`
}

// NewResetPasswordBody creates a new ResetPasswordBody for a ResetPassword.
func NewResetPasswordBody(resetToken, password string) ResetPasswordBody {
	resetPasswordRequest := ResetPasswordBody{
		ResetToken: resetToken,
		Password:   password,
	}

	return resetPasswordRequest
}
