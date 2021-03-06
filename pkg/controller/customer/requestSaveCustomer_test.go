// Copyright (c) 2021 Circutor S.A. All rights reserved.

package customer_test

import (
	"net/http"
	"testing"
	"time"

	dataCall "github.com/circutor/common-library/pkg/data"
	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/customer"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailBodyEncodeSaveCustomers(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	saveCustomerBody := core.NewSaveCustomerBody(
		"CUSTOMER",
		"00000000-0000-0000-0000-000000000000",
		"Spain",
		"Catalonia",
		"+34 123456789",
		"user@domine.com",
		"user@domine.com",
		"user@domine.com",
		"Circutor",
		"Aerospace",
		"CEO",
		"en_GB",
		time.Now().UnixNano()/int64(time.Millisecond))

	token := "Bearer token_value"

	data := new(dataMock.InterfaceDataMock)
	data.On("BodyEncode", saveCustomerBody).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := customer.NewControllerCustomerMock("/", "", "", data, nil)

	status, _, _ := controller.SaveCustomer(saveCustomerBody, token)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestLogin(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	saveCustomerBody := core.NewSaveCustomerBody(
		"CUSTOMER",
		"00000000-0000-0000-0000-000000000000",
		"Spain",
		"Catalonia",
		"+34 123456789",
		"user@domine.com",
		"user@domine.com",
		"user@domine.com",
		"Circutor",
		"Aerospace",
		"CEO",
		"en_GB",
		time.Now().UnixNano()/int64(time.Millisecond))

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(saveCustomerBody)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", saveCustomerBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	request.On("CreateNewRequest", "POST", "/api/customer", "Bearer token_value", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := customer.NewControllerCustomerMock("/", "", "", data, request)

	status, _, _ := controller.SaveCustomer(saveCustomerBody, "Bearer token_value")
	assert.Equal(t, http.StatusInternalServerError, status)
}
