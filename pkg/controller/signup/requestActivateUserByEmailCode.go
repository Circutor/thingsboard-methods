// Copyright (c) 2021 Circutor S.A. All rights reserved.

package signup

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
)

// ActivateUserByEmailCode activation user's account by email code.
func (c *ControllerSignUp) ActivateUserByEmailCode(emailCode string) (int, map[string]interface{}, error) {
	query := map[string]interface{}{
		"emailCode": emailCode,
	}

	url := c.TB.URLTBServer + noauth + activateUser

	resBody, status, err := request.CreateNewRequest(http.MethodPost, url, "", nil, query)
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

			return status, dataError, errors.NewErrFound(
				fmt.Sprint(thingsBoard), fmt.Sprint("ActivateUserByEmailCode ->", message))
		}
	}

	return status, nil, nil
}
