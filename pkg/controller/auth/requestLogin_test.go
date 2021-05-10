// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/circutor/thingsboard-methods/pkg/controller/auth"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginSuccess(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	userLogin := core.LoginBody{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	controller := auth.NewControllerAuth(cfg.URL, cfg.Username, cfg.Password)

	status, _, err := controller.Login(userLogin)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, status)
}

func TestLoginFailed(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	testCases := []struct {
		name     string
		body     core.LoginBody
		status   int
		respBody map[string]interface{}
	}{
		{
			name: "Authentication failed",
			body: core.LoginBody{
				Username: cfg.Username,
				Password: "",
			},
			status:   401,
			respBody: map[string]interface{}{"message": "Authentication failed"},
		},
		{
			name: "Invalid username or password",
			body: core.LoginBody{
				Username: cfg.Username,
				Password: "invalid_password",
			},
			status:   401,
			respBody: map[string]interface{}{"message": "Invalid username or password"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			respBodyByte, err := json.Marshal(tc.respBody)
			require.NoError(t, err)

			controller := auth.NewControllerAuth(cfg.URL, cfg.Username, cfg.Password)

			status, message, _ := controller.Login(tc.body)

			messageError, err := json.Marshal(message)
			require.NoError(t, err)

			assert.Equal(t, tc.status, status)
			assert.Equal(t, string(respBodyByte), string(messageError))
		})
	}
}
