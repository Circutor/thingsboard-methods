// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth_test

import (
	"encoding/json"
	"testing"

	"github.com/circutor/thingsboard-methods/pkg/controller/auth"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChangePasswordFailed(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	testCases := []struct {
		name     string
		status   int
		body     core.ChangePasswordBody
		respBody map[string]interface{}
	}{
		{
			name:     "Authentication failed",
			body:     core.NewChangePasswordBody("password_1", "password_2"),
			status:   401,
			respBody: map[string]interface{}{"message": "Authentication failed"},
		},
		{
			name:     "Authentication failed",
			body:     core.NewChangePasswordBody("", "password_2"),
			status:   401,
			respBody: map[string]interface{}{"message": "Authentication failed"},
		},
		{
			name:     "Authentication failed",
			body:     core.NewChangePasswordBody("password_1", ""),
			status:   401,
			respBody: map[string]interface{}{"message": "Authentication failed"},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			respBodyByte, err := json.Marshal(tc.respBody)
			require.NoError(t, err)

			controller := auth.NewControllerAuth(cfg.URL, cfg.Username, cfg.Password)

			status, message, _ := controller.ChangePassword(tc.body, "")

			messageError, err := json.Marshal(message)
			require.NoError(t, err)

			assert.Equal(t, tc.status, status)
			assert.Equal(t, string(respBodyByte), string(messageError))
		})
	}
}
