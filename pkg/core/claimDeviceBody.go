// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// ClaimDeviceBody body for request claimDevice.
type ClaimDeviceBody struct {
	SecretKey string `json:"secretKey"`
}

// NewClaimDeviceBody create a new ClaimDeviceBody for a claimDevice.
func NewClaimDeviceBody(secretKey string) ClaimDeviceBody {
	claimDevice := ClaimDeviceBody{
		SecretKey: secretKey,
	}

	return claimDevice
}
