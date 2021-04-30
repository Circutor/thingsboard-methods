// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

import "time"

// SaveUserBody body for request saveUser.
type SaveUserBody struct {
	UserID      EntityID `json:"id"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	Authority   string   `json:"authority"`
	CreatedTime int64    `json:"createdTime"`
	CustomerID  EntityID `json:"customerId"`
}

// NewSaveUserBody creates a new SaveUserBody for a saveUser.
func NewSaveUserBody(userID, userType, firstName, lastName, email, name, authority,
	customerID, customerType string) SaveUserBody {
	saveUserBody := SaveUserBody{
		UserID: EntityID{
			EntityType: userType,
			ID:         userID,
		},
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Name:        name,
		Authority:   authority,
		CreatedTime: time.Now().UnixNano() / int64(time.Millisecond),
		CustomerID: EntityID{
			EntityType: customerType,
			ID:         customerID,
		},
	}

	return saveUserBody
}
