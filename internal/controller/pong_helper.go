package controller

import (
	"common-go-example/internal/config"
	"github.com/kintohub/common-go/server/middleware"
	"github.com/valyala/fasthttp"
)

func (c *Controller) panicProcessPong(msg string) string {
	var resMsg string

	const disabledMessage = "pong is currently on vacation and cannot be found"

	if config.PongEnabled {
		if config.PongOverrideMessage == "" {
			resMsg = msg
		} else {
			resMsg = config.PongOverrideMessage
		}
	} else {
		middleware.PanicClientErrorWithMessage(fasthttp.StatusNotFound, disabledMessage)
	}

	return resMsg
}
