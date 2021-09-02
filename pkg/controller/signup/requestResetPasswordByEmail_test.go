// Copyright (c) 2021 Circutor S.A. All rights reserved.

package signup_test

import (
	"net/http"
	"testing"

	dataCall "github.com/circutor/common-library/pkg/data"
	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/signup"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailBodyEncodeResetPasswordByEmail(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	resetPasswordByEmailBody := core.NewResetPasswordByEmailBody("user@circutor.com")

	data := new(dataMock.InterfaceDataMock)
	data.On("BodyEncode", resetPasswordByEmailBody).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := signup.NewControllerSignUpMock("/", "", "", data, nil)

	status, _, _ := controller.ResetPasswordByEmail(resetPasswordByEmailBody, "Bearer token_value")
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestResetPasswordByEmail(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	resetPasswordByEmailBody := core.NewResetPasswordByEmailBody("user@circutor.com")

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(resetPasswordByEmailBody)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", resetPasswordByEmailBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/noauth/resetPasswordByEmail"

	request.On("CreateNewRequest", "POST", URL, "Bearer token_value", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := signup.NewControllerSignUpMock("/", "", "", data, request)

	status, _, _ := controller.ResetPasswordByEmail(resetPasswordByEmailBody, "Bearer token_value")
	assert.Equal(t, http.StatusInternalServerError, status)
}
