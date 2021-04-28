// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

type ThingsBoard struct {
	URLTBServer       string
	UserAdminName     string
	UserAdminPassword string
}

// New creates a new config ThingsBoard.
func NewThingsBoard(urlServer, userName, userPassword string) ThingsBoard {
	thingsBoardConfig := ThingsBoard{
		URLTBServer:       urlServer,
		UserAdminName:     userName,
		UserAdminPassword: userPassword,
	}

	return thingsBoardConfig
}

// EntityID struct id of type entity.
type EntityID struct {
	EntityType string `json:"entityType"`
	ID         string `json:"id"`
}
