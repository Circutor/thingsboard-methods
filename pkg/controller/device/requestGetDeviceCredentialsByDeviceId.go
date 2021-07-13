// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// GetDeviceCredentialsByDeviceID get credentials from device witch by ID.
func (c ControllerDevice) GetDeviceCredentialsByDeviceID(deviceID, token string) (int, map[string]interface{}, error) {
	resBody, status, err := c.Request.CreateNewRequest(
		http.MethodGet, c.TB.URLTBServer+device+"/"+deviceID+credentials, token, nil, nil)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	responseBody, err := c.Data.BodyDecodeToMap(resBody)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	if message, ok := responseBody["message"]; ok {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(fmt.Sprint(message)))

		return status, dataError, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetDeviceCredentialsByDeviceID ->", message))
	}

	return status, responseBody, nil
}
