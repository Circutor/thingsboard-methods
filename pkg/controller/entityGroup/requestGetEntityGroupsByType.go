// Copyright (c) 2021 Circutor S.A. All rights reserved.

package entitygroup

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
	"github.com/circutor/thingsboard-methods/pkg/data"
)

// GetEntityGroupsByType get all entities from group type.
func (c *ControllerEntityGroup) GetEntityGroupsByType(groupType, token string) (int, []interface{}, error) {
	url := c.TB.URLTBServer + groups + "/" + groupType

	resBody, status, err := request.CreateNewRequest(http.MethodGet, url, token, nil, nil)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(string(resBody)))

		return status, dataError, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("GetEntityGroupsByType ->", string(resBody)))
	}

	responseBody, err := data.BodyDecode(resBody)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	return status, responseBody, nil
}
