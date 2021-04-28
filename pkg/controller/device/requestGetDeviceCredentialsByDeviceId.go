// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
)

// GetDeviceCredentialsByDeviceID get credentials from device witch by ID.
func (c ControllerDevice) GetDeviceCredentialsByDeviceID(deviceID, token string) (int, map[string]interface{}, error) {
	resBody, status, err := request.CreateNewRequest(
		http.MethodGet, c.TB.URLTBServer+device+"/"+deviceID+credentials, token, nil, nil)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	responseBody, err := data.BodyDecode(resBody)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	if message, ok := responseBody["message"]; ok {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(fmt.Sprint(message)))

		return status, dataError, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetDeviceCredentialsByDeviceID ->", message))
	}

	return status, responseBody, nil
}
