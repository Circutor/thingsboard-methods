// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// GetAttributeKeysByScope get keys from entity from by scope.
func (c *ControllerTelemetry) GetAttributeKeysByScope(entityType, entityID, scope,
	token string) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + getAttributesKeys + "/" + scope

	resBody, status, err := c.Request.CreateNewRequest(http.MethodGet, url, token, nil, nil)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if status == http.StatusUnauthorized {
		responseBody, _ := c.Data.BodyDecodeToMap(resBody)

		return status, map[string]interface{}{"message": responseBody["message"]}, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetAttributeKeysByScope ->", responseBody["message"]))
	}

	if status == http.StatusForbidden || status == http.StatusNotFound {
		return status, map[string]interface{}{"message": string(resBody)}, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetAttributeKeysByScope ->", string(resBody)))
	}

	responseBody, err := c.Data.BodyDecodeToArray(resBody)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToArray(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, map[string]interface{}{"message": dataError[0]}, fmt.Errorf("%w", err)
	}

	return status, map[string]interface{}{"attributes": responseBody}, nil
}
