package model

import (
	"common-go-example/internal/config"
	"common-go-example/internal/types"
	"errors"
)

func Pong(request types.PingRequest) (*types.PingResponse, error) {
	response := types.PingResponse{}

	const disabledMessage = "pong is currently on vacation and cannot be found"

	if config.PongEnabled {
		if config.PongOverrideMessage == "" {
			response.Message = request.Message
		} else {
			response.Message = config.PongOverrideMessage
		}
	} else {
		return nil, errors.New(disabledMessage)
	}

	return &response, nil
}
