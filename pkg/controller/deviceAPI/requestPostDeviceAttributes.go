// Copyright (c) 2021 Circutor S.A. All rights reserved.

//nolint:dupl
package deviceapi

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
)

// PostDeviceAttributes save attributes in ThingsBoard to client scope.
func (c *ControllerDeviceAPI) PostDeviceAttributes(deviceToken string,
	attrBody interface{}) (int, map[string]interface{}, error) {
	body, err := data.BodyEncode(attrBody)
	if err != nil {
		return http.StatusInternalServerError, nil, fmt.Errorf("%w", err)
	}

	url := c.TB.URLTBServer + deviceAPI + "/" + deviceToken + attributes

	resBody, status, err := request.CreateNewRequest(http.MethodPost, url, "", body, nil)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		responseBody, err := data.BodyDecode(resBody)
		if err != nil {
			dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

			return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
		}

		if message, ok := responseBody["message"]; ok {
			dataError, _ := data.ResponseDecode(errors.NewErrMessage(fmt.Sprint(message)))

			return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("PostDeviceAttributes ->", message))
		}
	}

	return status, nil, nil
}
