// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

func (c *ControllerTelemetry) DeleteEntityAttributes(entityType, entityID, scope, token string,
	query map[string]interface{}) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + "/" + scope

	resBody, status, err := c.Request.CreateNewRequest(http.MethodDelete, url, token, nil, query)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
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
				fmt.Sprint(thingsBoard), fmt.Sprint("DeleteEntityAttributes ->", message))
		}
	}

	return status, nil, nil
}
