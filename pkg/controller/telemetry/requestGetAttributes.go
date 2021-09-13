// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// GetAttributes get attributes value from entity.
func (c *ControllerTelemetry) GetAttributes(entityType, entityID, token string,
	query map[string]interface{}) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + getAttributesValues

	resBody, status, err := c.Request.CreateNewRequest(http.MethodGet, url, token, nil, query)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(err.Error()))

		return status, map[string]interface{}{"message": dataError[0]}, fmt.Errorf("%w", err)
	}

	if status == http.StatusUnauthorized {
		responseBody, _ := c.Data.BodyDecodeToMap(resBody)

		return status, map[string]interface{}{"message": responseBody["message"]}, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetAttributes ->", responseBody["message"]))
	}

	if status == http.StatusForbidden || status == http.StatusNotFound {
		return status, map[string]interface{}{"message": string(resBody)}, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetAttributes ->", string(resBody)))
	}

	responseBody, err := c.Data.BodyDecodeToArray(resBody)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, map[string]interface{}{"message": dataError[0]}, fmt.Errorf("%w", err)
	}

	return status, map[string]interface{}{"attributes": responseBody}, nil
}
