// Copyright (c) 2021 Circutor S.A. All rights reserved.

package auth_test

import (
	"github.com/circutor/thingsboard-methods/internal/config"
	"github.com/kelseyhightower/envconfig"
)

//nolint:gochecknoglobals
var cfg config.Config

//nolint:gochecknoinits
func init() {
	_ = envconfig.Process("tb", &cfg)
}
