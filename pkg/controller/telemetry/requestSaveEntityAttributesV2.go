// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/errors"
)

// SaveEntityAttributesV2 create entity attributes witch scope(server, hared or client).
func (c *ControllerTelemetry) SaveEntityAttributesV2(entityType, entityID, scope, token string,
	attributesBody map[string]interface{}) (int, map[string]interface{}, error) {
	body, err := c.Data.BodyEncode(attributesBody)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
	}

	url := c.TB.URLTBServer + telemetry + entityType + "/" + entityID + saveAttributes + scope

	resBody, status, err := c.Request.CreateNewRequest(http.MethodPost, url, token, body, nil)
	if err != nil {
		dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return status, dataError, fmt.Errorf("%w", err)
	}

	if !(status == http.StatusOK || status == http.StatusCreated) {
		responseBody, err := c.Data.BodyDecodeToMap(resBody)
		if err != nil {
			dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

			return http.StatusInternalServerError, dataError, fmt.Errorf("%w", err)
		}

		if message, ok := responseBody["message"]; ok {
			dataError, _ := c.Data.ResponseDecodeToMap(errors.NewErrMessage(fmt.Sprint(message)))

			return status, dataError, errors.NewErrFound(
				fmt.Sprint(thingsBoard), fmt.Sprint("SaveEntityAttributesV2 ->", message))
		}
	}

	return status, nil, nil
}
