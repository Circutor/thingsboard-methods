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

func TestFailBodyEncodeResetPassword(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	resetPasswordBody := core.NewResetPasswordBody("resetToken", "password")

	data := new(dataMock.InterfaceDataMock)
	data.On("BodyEncode", resetPasswordBody).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := signup.NewControllerSignUpMock("/", "", "", data, nil)

	status, _, _ := controller.ResetPassword(resetPasswordBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestResetPassword(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	resetPasswordBody := core.NewResetPasswordBody("resetToken", "password")

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(resetPasswordBody)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", resetPasswordBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/noauth/resetPassword"

	request.On("CreateNewRequest", "POST", URL, "", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := signup.NewControllerSignUpMock("/", "", "", data, request)

	status, _, _ := controller.ResetPassword(resetPasswordBody)
	assert.Equal(t, http.StatusInternalServerError, status)
}
