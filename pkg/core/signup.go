// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// SignUpBody body for a signUp body.
type SignUpBody struct {
	Email             string `json:"email"`
	FistName          string `json:"firstName"`
	LastName          string `json:"lastName"`
	Password          string `json:"password"`
	RecaptchaResponse string `json:"recaptchaResponse"`
}

// NewSignUpBody create a new SignUpBody for a SignUp.
func NewSignUpBody(email, fistName, lastName, password, recaptchaResponse string) SignUpBody {
	signUpReBody := SignUpBody{
		Email:             email,
		FistName:          fistName,
		LastName:          lastName,
		Password:          password,
		RecaptchaResponse: recaptchaResponse,
	}

	return signUpReBody
}
