// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// ChangePasswordBody body for request change password.
type ChangePasswordBody struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// NewChangePasswordBody create a new ChangePasswordBody for a change password.
func NewChangePasswordBody(currentPassword, newPassword string) ChangePasswordBody {
	changePasswordBody := ChangePasswordBody{
		CurrentPassword: currentPassword,
		NewPassword:     newPassword,
	}

	return changePasswordBody
}
