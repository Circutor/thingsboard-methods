// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// LoginBody body for request login.
type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewLoginBody create a new LoginBody for a Login.
func NewLoginBody(username, password string) LoginBody {
	login := LoginBody{
		Username: username,
		Password: password,
	}

	return login
}
