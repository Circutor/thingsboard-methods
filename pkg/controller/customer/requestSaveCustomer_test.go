// Copyright (c) 2021 Circutor S.A. All rights reserved.

package customer_test

import (
	"net/http"
	"testing"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
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
		"00180dfc-98f8-11eb-a8b3-0242ac130003",
		"Spain",
		"Catalonia",
		"+34 123456789",
		"user@domine.com",
		"user@domine.com",
		"user@domine.com",
		"Circutor",
		"Aerospace",
		"CEO", "en_GB")

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
