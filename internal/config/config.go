package config

import (
	"github.com/kintohub/common-go/utils/env"
	"strconv"
)

var (
	PongOverrideMessage = env.Get("PONG_OVERRIDE_MESSAGE", "")
	PongEnabled         bool
)

func init() {
	enabled, err := strconv.ParseBool(env.Get("PONG_ENABLED", "false"))

	if err == nil {
		PongEnabled = enabled
	} else {
		panic("Invalid bool (true/false) string format for PONG_ENABLED env var")
	}
}
