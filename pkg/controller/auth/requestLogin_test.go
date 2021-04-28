// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"thingsboard-methods/pkg/controller/auth"
	"thingsboard-methods/pkg/core"

	"github.com/jbrodriguez/mlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginTB(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	userLogin := core.LoginBody{
		Username: "",
		Password: "",
	}

	controller := auth.NewControllerAuth("", "", "")

	status, _, _ := controller.Login(userLogin)

	assert.Equal(t, http.StatusOK, status)
}

func TestLoginErrorTB(t *testing.T) {
	t.Parallel()

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	userLogin := core.LoginBody{
		Username: "",
		Password: "",
	}

	respBody := map[string]interface{}{
		"message": "Invalid username or password",
	}

	respBodyByte, err := json.Marshal(respBody)
	require.NoError(t, err)

	controller := auth.NewControllerAuth("", "", "")

	status, message, err := controller.Login(userLogin)
	require.NoError(t, err)

	messageError, err := json.Marshal(message)
	require.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, status)
	assert.Equal(t, string(respBodyByte), string(messageError))
}
