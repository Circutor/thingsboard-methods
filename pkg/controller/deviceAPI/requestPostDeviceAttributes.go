// Copyright (c) 2021 Circutor S.A. All rights reserved.

//nolint:dupl
package deviceapi

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// PostDeviceAttributes save attributes in ThingsBoard to client scope.
func (c *ControllerDeviceAPI) PostDeviceAttributes(deviceToken string,
	attrBody interface{}) (int, map[string]interface{}, error) {
	body, err := c.Data.BodyEncode(attrBody)
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("%w", err)
	}

	url := c.TB.URLTBServer + deviceAPI + "/" + deviceToken + attributes

	resBody, status, err := c.Request.CreateNewRequest(http.MethodPost, url, "", body, nil)
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

			return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("PostDeviceAttributes ->", message))
		}
	}

	return status, nil, nil
}
