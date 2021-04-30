// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
	"github.com/circutor/thingsboard-methods/pkg/core"
)

func (c *ControllerDevice) ClaimDevice(deviceName, token string,
	claimDeviceBody core.ClaimDeviceBody) (int, map[string]interface{}, error) {
	body, err := data.BodyEncode(claimDeviceBody)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	url := c.TB.URLTBServer + claimDevice + deviceName + "/" + claim

	resBody, status, err := request.CreateNewRequest(http.MethodPost, url, token, body, nil)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		responseBody, err := data.BodyDecode(resBody)
		if err != nil {
			dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

			return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
		}

		if message, ok := responseBody["message"]; ok {
			dataError, _ := data.ResponseDecode(errors.NewErrMessage(fmt.Sprint(message)))

			return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("ClaimDevice ->", message))
		}
	}

	return status, nil, nil
}
