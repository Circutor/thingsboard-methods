// Copyright (c) 2021 Circutor S.A. All rights reserved.

//nolint:dupl
package auth

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
	"github.com/circutor/thingsboard-methods/pkg/core"
)

func (c ControllerAuth) Login(loginBody core.LoginBody) (int, map[string]interface{}, error) {
	body, err := data.BodyEncode(loginBody)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	url := c.TB.URLTBServer + auth + login

	resBody, status, err := request.CreateNewRequest(http.MethodPost, url, "", body, nil)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	responseBody, err := data.BodyDecode(resBody)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	if message, ok := responseBody["message"]; ok {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(fmt.Sprint(message)))

		return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("Login ->", message))
	}

	return status, responseBody, nil
}
