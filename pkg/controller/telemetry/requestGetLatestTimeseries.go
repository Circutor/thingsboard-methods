// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
	dataTB "github.com/circutor/thingsboard-methods/pkg/data"
)

// GetLatestTimeseries get values last time series.
func (c *ControllerTelemetry) GetLatestTimeseries(entityType, entityID, token string,
	query map[string]interface{}) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + getTimeseriesValues

	resBody, status, err := request.CreateNewRequest(http.MethodGet, url, token, nil, query)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		dataError, _ := dataTB.ResponseDecode(errors.NewErrMessage(string(resBody)))

		message := map[string]interface{}{
			"message": dataError[0].(string),
		}

		return status, message, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetLatestTimeseries ->", string(resBody)))
	}

	responseBody, err := data.BodyDecode(resBody)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	return status, responseBody, nil
}
