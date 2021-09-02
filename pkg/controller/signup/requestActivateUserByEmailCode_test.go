// Copyright (c) 2021 Circutor S.A. All rights reserved.

package signup_test

import (
	"net/http"
	"testing"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/signup"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailRequestActivateUserByEmailCode(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	query := map[string]interface{}{
		"emailCode": "emailCode",
	}
	data := new(dataMock.InterfaceDataMock)

	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/noauth/activateByEmailCode"

	request.On("CreateNewRequest", "POST", URL, "", nil, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := signup.NewControllerSignUpMock("/", "", "", data, request)

	status, _, _ := controller.ActivateUserByEmailCode("emailCode")
	assert.Equal(t, http.StatusInternalServerError, status)
}
