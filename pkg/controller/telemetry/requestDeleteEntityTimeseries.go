// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// DeleteEntityTimeseries delete timeSeries for a device.
func (c *ControllerTelemetry) DeleteEntityTimeseries(entityType, entityID, token string,
	query map[string]interface{}) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + getTimeseriesDelete

	resBody, status, err := c.Request.CreateNewRequest(http.MethodDelete, url, token, nil, query)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if status == http.StatusUnauthorized {
		responseBody, _ := c.Data.BodyDecodeToMap(resBody)

		return status, map[string]interface{}{"message": responseBody["message"]}, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("DeleteEntityTimeseries ->", responseBody["message"]))
	}

	if status == http.StatusForbidden || status == http.StatusNotFound || status == http.StatusBadRequest {
		return status, map[string]interface{}{"message": string(resBody)}, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("DeleteEntityTimeseries ->", string(resBody)))
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		responseBody, err := c.Data.BodyDecodeToMap(resBody)
		if err != nil {
			dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

			return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
		}

		if message, ok := responseBody["message"]; ok {
			dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(fmt.Sprint(message)))

			return status, dataError, errors.NewErrFound(
				fmt.Sprint(thingsBoard), fmt.Sprint("DeleteEntityTimeseries ->", message))
		}
	}

	return status, nil, nil
}
