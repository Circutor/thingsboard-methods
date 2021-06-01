// Copyright (c) 2021 Circutor S.A. All rights reserved.

//nolint:dupl
package device

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
)

// GetCustomerDevices get all devices from Customer User.
func (c *ControllerDevice) GetCustomerDevices(customerID, token string,
	query map[string]interface{}) (int, map[string]interface{}, error) {
	url := c.TB.URLTBServer + devicesCustomer + customerID + devices

	resBody, status, err := request.CreateNewRequest(http.MethodGet, url, token, nil, query)
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

		return status, dataError, errors.NewErrFound(fmt.Sprint(thingsBoard), fmt.Sprint("GetCustomerDevices ->", message))
	}

	return status, responseBody, nil
}
