// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/request"
	"github.com/circutor/thingsboard-methods/pkg/controller/device"
	"github.com/circutor/thingsboard-methods/pkg/controller/telemetry"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAttributesKeysByScopeSuccess(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	controller := telemetry.NewControllerTelemetry(cfg.URL, cfg.Username, cfg.Password)

	// create map and buffer byte to send createDevice.
	createDevice, b := map[string]interface{}{"serialId": "00000000000000", "type": "DeviceTest"}, new(bytes.Buffer)
	_ = json.NewEncoder(b).Encode(createDevice)

	resBody, _, _ := request.CreateNewRequest(http.MethodPost, manufacture, "", b, nil)
	deviceResponse, _ := data.BodyDecode(resBody)

	id := fmt.Sprintf("%v", deviceResponse["id"])

	fmt.Println(id)

	status, _, err := controller.GetAttributeKeysByScope("DEVICE", id, "CLIENT_SCOPE", token)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, status)

	controllerDevice := device.NewControllerDevice(cfg.URL, cfg.Username, cfg.Password)
	status, _, err = controllerDevice.DeleteDevice(fmt.Sprintf("%v", deviceResponse["id"]), token)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, status)
}

func TestGetAttributesKeysByScopeDeviceNotExist(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	controller := telemetry.NewControllerTelemetry(cfg.URL, cfg.Username, cfg.Password)

	status, _, _ := controller.GetAttributeKeysByScope(
		"DEVICE", "00000000-0000-0000-0000-000000000000", "CLIENT_SCOPE", token)

	assert.Equal(t, http.StatusNotFound, status)
}
