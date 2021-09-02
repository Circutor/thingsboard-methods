// Copyright (c) 2021 Circutor S.A. All rights reserved.

package entitygroup_test

import (
	"net/http"
	"testing"

	dataMock "github.com/circutor/common-library/pkg/data/mocks"
	"github.com/circutor/common-library/pkg/errors"
	requestMock "github.com/circutor/common-library/pkg/request/mocks"
	entitygroup "github.com/circutor/thingsboard-methods/pkg/controller/entityGroup"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
)

func TestFailRequestGetEntityGroupsByType(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	groupType := "DEVICE"

	query := map[string]interface{}(nil)
	data := new(dataMock.InterfaceDataMock)

	data.On("ResponseDecodeToMap", errors.NewErrMessage("error in create request")).
		Return(map[string]interface{}{"message": "error in encode request body"}, nil)

	request := new(requestMock.InterfaceRequestMock)

	URL := "/api/entityGroups/" + groupType

	request.On("CreateNewRequest", "GET", URL, "Bearer token_value", nil, query).
		Return(nil, 500, errors.NewErrFound("error in create request"))

	controller := entitygroup.NewControllerEntityGroupMock("/", "", "", data, request)

	status, _, _ := controller.GetEntityGroupsByType(groupType, "Bearer token_value")
	assert.Equal(t, http.StatusInternalServerError, status)
}
