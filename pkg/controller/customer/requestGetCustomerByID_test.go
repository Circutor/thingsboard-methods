// Copyright (c) 2021 Circutor S.A. All rights reserved.

package customer_test

import (
	"net/http"
	"testing"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/customer"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailRequestGetCustomerByID(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	deviceID := "00000000-0000-0000-0000-000000000000-6a"

	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	request.On("CreateNewRequest", "GET", "/api/customer"+"/"+deviceID, "Bearer token_value", nil, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := customer.NewControllerCustomerMock("/", "", "", data, request)

	status, _, _ := controller.GetCustomerByID(deviceID, "Bearer token_value")
	assert.Equal(t, http.StatusInternalServerError, status)
}
