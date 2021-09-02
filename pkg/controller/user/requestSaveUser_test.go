// Copyright (c) 2021 Circutor S.A. All rights reserved.

package user_test

import (
	"net/http"
	"testing"
	"time"

	dataCall "github.com/circutor/common-library/pkg/data"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/user"
	"github.com/circutor/thingsboard-methods/pkg/core"

	"github.com/circutor/common-library/pkg/errors"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailBodyEncodeSaveUser(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	saveUserBody := core.NewSaveUserBody(
		"00180dfc-98f8-11eb-a8b3-0242ac130003",
		"USER",
		"Lorem",
		"Ipsum",
		"user@circutor.com",
		"user",
		"TENANT_ADMIN",
		"00180dfc-98f8-11eb-a8b3-0242ac130003",
		"USER",
		time.Now().UnixNano()/int64(time.Millisecond))

	data := new(dataMock.InterfaceDataMock)
	data.On("BodyEncode", saveUserBody).
		Return(nil, errors.NewErrFound("error in encode request body"))
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in encode request body")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	controller := user.NewControllerUserMock("/", "", "", data, nil)

	status, _, _ := controller.SaveUser(saveUserBody, "Bearer token_value")
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestFailRequestSaveUser(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	saveUserBody := core.NewSaveUserBody(
		"00000000-0000-0000-0000-000000000000",
		"USER",
		"Lorem",
		"Ipsum",
		"user@circutor.com",
		"user",
		"TENANT_ADMIN",
		"00000000-0000-0000-0000-000000000000",
		"USER",
		time.Now().UnixNano()/int64(time.Millisecond))

	d := dataCall.NewData()
	respBody, _ := d.BodyEncode(saveUserBody)
	query := map[string]interface{}(nil)

	data := new(dataMock.InterfaceDataMock)

	data.On("BodyEncode", saveUserBody).
		Return(respBody, nil)
	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	request.On("CreateNewRequest", "POST", "/api/user", "Bearer token_value", respBody, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := user.NewControllerUserMock("/", "", "", data, request)

	status, _, _ := controller.SaveUser(saveUserBody, "Bearer token_value")
	assert.Equal(t, http.StatusInternalServerError, status)
}
