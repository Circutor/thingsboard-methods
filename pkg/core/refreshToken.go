// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// RefreshTokenBody body for request login.
type RefreshTokenBody struct {
	RefreshToken string `json:"refreshToken"`
}

// NewRefreshTokenBody creates a new RefreshTokenBody for a refreshToken.
func NewRefreshTokenBody(refreshToken string) RefreshTokenBody {
	refreshTokenBody := RefreshTokenBody{
		RefreshToken: refreshToken,
	}

	return refreshTokenBody
}
