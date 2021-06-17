// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// Logout makes a logout User Tenant.
func (c *ControllerAuth) Logout(token string) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + auth + logout

	resBody, status, err := c.Request.CreateNewRequest(http.MethodPost, url, token, nil, nil)
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

			return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("Logout ->", message))
		}
	}

	return status, nil, nil
}
