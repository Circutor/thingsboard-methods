// Copyright (c) 2021 Circutor S.A. All rights reserved.

//nolint:dupl
package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// GetAttributes get attributes value from entity.
func (c *ControllerTelemetry) GetAttributes(entityType, entityID,
	token string, query map[string]interface{}) (int, []interface{}, error) {
	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + getAttributesValues

	resBody, status, err := c.Request.CreateNewRequest(http.MethodGet, url, token, nil, query)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(string(resBody)))

		return status, dataError, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetAttributes ->", string(resBody)))
	}

	responseBody, err := c.Data.BodyDecodeToArray(resBody)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	return status, responseBody, nil
}
