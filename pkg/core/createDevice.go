// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

import (
	"strings"
	"time"
)

// CreateDeviceBody body for request createDevice.
type CreateDeviceBody struct {
	CreatedTime int64  `json:"createdTime"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

// NewCreateDeviceBody create a new CreateDeviceBody for a CreateDevice.
func NewCreateDeviceBody(serialID, deviceType string) CreateDeviceBody {
	createDeviceBody := CreateDeviceBody{
		CreatedTime: time.Now().UnixNano() / int64(time.Millisecond),
		Label:       strings.ToUpper(deviceType),
		Name:        serialID + "_" + strings.ToUpper(deviceType),
		Type:        strings.ToUpper(deviceType),
	}

	return createDeviceBody
}
