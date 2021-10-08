// Copyright (c) 2021 Circutor S.A. All rights reserved.

package user

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// GetUserByID get user by userID info.
func (c *ControllerUser) GetUserByID(userID, token string) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + user + "/" + userID

	resBody, status, err := c.Request.CreateNewRequest(http.MethodGet, url, token, nil, nil)
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

		return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("GetUserUsers ->", message))
	}

	return status, responseBody, nil
}
