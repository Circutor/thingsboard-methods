// Copyright (c) 2021 Circutor S.A. All rights reserved.

package deviceapi

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
)

func (c *ControllerDeviceAPI) PostTelemetry(deviceToken string,
	query map[string]interface{}) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + deviceAPI + "/" + deviceToken + telemetry

	resBody, status, err := request.CreateNewRequest(http.MethodPost, url, "", nil, query)
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

			return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("PostTelemetry ->", message))
		}
	}

	return status, nil, nil
}
