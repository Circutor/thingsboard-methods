// Copyright (c) 2021 Circutor S.A. All rights reserved.

package deviceapi

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// GetDeviceAttributes get attributes in ThingsBoard to client scope.
func (c *ControllerDeviceAPI) GetDeviceAttributes(deviceToken string,
	query map[string]interface{}) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + deviceAPI + "/" + deviceToken + attributes

	resBody, status, err := c.Request.CreateNewRequest(http.MethodGet, url, "", nil, query)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	responseBody, err := c.Data.BodyDecodeToMap(resBody)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	if message, ok := responseBody["message"]; ok {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(fmt.Sprint(message)))

		return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("GetDeviceAttributes ->", message))
	}

	return status, responseBody, nil
}
