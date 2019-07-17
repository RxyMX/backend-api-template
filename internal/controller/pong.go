package controller

import (
	"common-go-example/internal/config"
	"errors"
	"github.com/kintohub/common-go/server/middleware"
	"github.com/valyala/fasthttp"
)

func pong(request PingRequest) *PingResponse {
	response := PingResponse{}

	const disabledMessage = "pong is currently on vacation and cannot be found"

	if config.PongEnabled {
		if config.PongOverrideMessage == "" {
			response.Message = request.Message
		} else {
			response.Message = config.PongOverrideMessage
		}
	} else {
		middleware.PanicClientErrorWithMessage(fasthttp.StatusNotFound,
			disabledMessage,
			errors.New(disabledMessage))
	}

	return &response
}
