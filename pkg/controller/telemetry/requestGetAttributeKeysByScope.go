// Copyright (c) 2021 Circutor S.A. All rights reserved.

//nolint:dupl
package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
	"github.com/circutor/thingsboard-methods/internal/data"
)

// GetAttributeKeysByScope get keys from entity from by scope.
func (c *ControllerTelemetry) GetAttributeKeysByScope(entityType, entityID, scope,
	token string) (int, []interface{}, error) {
	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + getAttributesKeys + "/" + scope

	resBody, status, err := request.CreateNewRequest(http.MethodGet, url, token, nil, nil)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(string(resBody)))

		return status, dataError, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetAttributeKeysByScope ->", string(resBody)))
	}

	responseBody, err := data.BodyDecode(resBody)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	return status, responseBody, nil
}
