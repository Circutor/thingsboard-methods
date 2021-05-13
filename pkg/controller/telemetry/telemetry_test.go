// Copyright (c) 2021 Circutor S.A. All rights reserved.

package telemetry_test

import (
	"fmt"

	"github.com/circutor/thingsboard-methods/internal/config"
	"github.com/circutor/thingsboard-methods/pkg/controller/auth"
	"github.com/circutor/thingsboard-methods/pkg/core"
	"github.com/jbrodriguez/mlog"
	"github.com/kelseyhightower/envconfig"
)

//nolint:gochecknoglobals
var (
	cfg         config.Config
	manufacture string
	token       string
)

//nolint:gochecknoinits
func init() {
	_ = envconfig.Process("tb", &cfg)
	manufacture = "http://0.0.0.0:60010/api/v1/manufacture-service/createDevice"

	mlog.StartEx(mlog.LevelTrace, "", 0, 0)

	// get authorization admin
	controllerAuth := auth.NewControllerAuth(cfg.URL, cfg.Username, cfg.Password)

	loginBody := core.NewLoginBody(cfg.Username, cfg.Password)
	_, authorization, _ := controllerAuth.Login(loginBody)

	token = fmt.Sprintf("Bearer %v", authorization["token"])
}
