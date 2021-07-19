// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth_test

import (
	"net/http"
	"testing"

	dataCall "github.com/circutor/common-library/pkg/data"
	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/auth"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailBodyEncodeRefreshToken(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	body := core.NewRefreshTokenBody("refreshToken_value")

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", body).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := auth.NewControllerAuthMock("/", "", "", data, nil)

	status, _, _ := controller.RefreshToken(body)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestRefreshToken(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	body := core.NewRefreshTokenBody("refreshToken_value")

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(body)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", body).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	request.On("CreateNewRequest", "POST", "/api/auth/token", "", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := auth.NewControllerAuthMock("/", "", "", data, request)

	status, _, _ := controller.RefreshToken(body)
	assert.Equal(t, http.StatusInternalServerError, status)
}
