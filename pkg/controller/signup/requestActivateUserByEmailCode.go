// Copyright (c) 2021 Circutor S.A. All rights reserved.

package signup

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// ActivateUserByEmailCode activation user's account by email code.
func (c *ControllerSignUp) ActivateUserByEmailCode(emailCode string) (int, map[string]interface{}, error) {
	query := map[string]interface{}{
		"emailCode": emailCode,
	}

	url := c.TB.URLTBServer + noauth + activateUser

	resBody, status, err := c.Request.CreateNewRequest(http.MethodPost, url, "", nil, query)
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
				fmt.Sprint(thingsBoard), fmt.Sprint("ActivateUserByEmailCode ->", message))
		}
	}

	return status, nil, nil
}
