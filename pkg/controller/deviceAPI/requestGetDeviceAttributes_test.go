// Copyright (c) 2021 Circutor S.A. All rights reserved.

package deviceapi_test

import (
	"net/http"
	"testing"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	deviceapi "github.com/circutor/thingsboard-methods/pkg/controller/deviceAPI"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailRequestGetDeviceAttributes(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	accessToken := "00000000000000000000"

	query := map[string]interface{}(nil)
	data := new(dataMock.InterfaceDataMock)

	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/v1/" + accessToken + "/attributes"

	request.On("CreateNewRequest", "GET", URL, "", nil, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := deviceapi.NewControllerDeviceAPIMock("/", "", "", data, request)

	status, _, _ := controller.GetDeviceAttributes(accessToken, query)
	assert.Equal(t, http.StatusInternalServerError, status)
}
