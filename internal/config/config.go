package config

import (
	"github.com/kintohub/utils-go/utils"
	"strconv"
)

var (
	LogLevel            = utils.Get("LOG_LEVEL", "info")
	ServerPort          = utils.Get("SERVER_PORT", "8080")
	PongOverrideMessage = utils.Get("PONG_OVERRIDE_MESSAGE", "")
	PongEnabled         bool
)

func init() {
	enabled, err := strconv.ParseBool(utils.Get("PONG_ENABLED", "false"))

	if err == nil {
		PongEnabled = enabled
	} else {
		panic("Invalid bool (true/false) string format for PONG_ENABLED env var")
	}
}
