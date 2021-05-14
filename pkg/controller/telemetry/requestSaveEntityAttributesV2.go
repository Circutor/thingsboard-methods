// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"fmt"
	"net/http"

	common "github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/request"
	"github.com/circutor/thingsboard-methods/pkg/data"
)

// SaveEntityAttributesV2  create entity attributes witch scope(server, hared or client).
func (c *ControllerTelemetry) SaveEntityAttributesV2(entityType, entityID, scope, token string,
	attributesBody map[string]interface{}) (int, []interface{}, error) {
	body, err := common.BodyEncode(attributesBody)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + saveAttributes + scope

	resBody, status, err := request.CreateNewRequest(http.MethodPost, url, token, body, nil)
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(string(resBody)))

		return status, dataError, errors.NewErrFound(
			fmt.Sprint(thingsBoard), fmt.Sprint("SaveEntityAttributesV2 ->", string(resBody)))
	}

	return status, nil, nil
}
