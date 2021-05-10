// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/circutor/thingsboard-methods/pkg/controller/auth"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogoutSuccess(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	userLogin := core.LoginBody{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	controller := auth.NewControllerAuth(cfg.URL, cfg.Username, cfg.Password)

	_, authorization, err := controller.Login(userLogin)
	require.NoError(t, err)

	status, _, err := controller.Logout(fmt.Sprintf("Bearer %v", authorization["token"]))
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, status)
}

func TestLogoutFailed(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	testCases := []struct {
		name     string
		status   int
		token    string
		respBody map[string]interface{}
	}{
		{
			name:     "Authentication failed",
			token:    "",
			status:   401,
			respBody: map[string]interface{}{"message": "Authentication failed"},
		},
		{
			name:     "Invalid username or password",
			token:    "Bearer ",
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

			status, message, _ := controller.Logout(fmt.Sprintf("Bearer %v", tc.token))

			messageError, err := json.Marshal(message)
			require.NoError(t, err)

			assert.Equal(t, tc.status, status)
			assert.Equal(t, string(respBodyByte), string(messageError))
		})
	}
}
