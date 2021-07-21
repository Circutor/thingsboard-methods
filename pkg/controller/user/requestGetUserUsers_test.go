// Copyright (c) 2021 Circutor S.A. All rights reserved.

package user_test

import (
	"net/http"
	"testing"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/thingsboard-methods/pkg/controller/user"

	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailRequestGetUserUsers(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	query := map[string]interface{}{"pageSize": 10, "page": 0}

	data := new(dataMock.InterfaceDataMock)

	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	request.On("CreateNewRequest", "GET", "/api/user/users", "Bearer token_value", nil, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := user.NewControllerUserMock("/", "", "", data, request)

	status, _, _ := controller.GetUserUsers("Bearer token_value", query)
	assert.Equal(t, http.StatusInternalServerError, status)
}
