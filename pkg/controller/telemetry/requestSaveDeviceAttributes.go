// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// SaveDeviceAttributes create device attributes witch scope(server or shared).
func (c *ControllerTelemetry) SaveDeviceAttributes(deviceID, scope, token string,
	attributesBody map[string]interface{}) (int, []interface{}, error) {
	body, err := c.Data.BodyEncode(attributesBody)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	url := c.TB.URLTBServer + telemetry + deviceID + "/" + scope

	resBody, status, err := c.Request.CreateNewRequest(http.MethodPost, url, token, body, nil)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(string(resBody)))

		return status, dataError, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("SaveDeviceAttributes ->", string(resBody)))
	}

	return status, nil, nil
}
