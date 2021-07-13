// Copyright (c) 2021 Circutor S.A. All rights reserved.

package customer

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

func (c *ControllerCustomer) GetTenantCustomer(email, token string) (int, map[string]interface{}, error) {
	query := map[string]interface{}{
		"customerTitle": email,
	}

	url := c.TB.URLTBServer + getTenantCustomer

	resBody, status, err := c.Request.CreateNewRequest(http.MethodGet, url, token, nil, query)
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

		return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("GetTenantCustomer ->", message))
	}

	return status, responseBody, nil
}
