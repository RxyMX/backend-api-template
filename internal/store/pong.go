package store

import (
	"common-go-example/internal/config"
	"common-go-example/internal/model"
	"errors"
)

func Pong(request model.PingRequest) (*model.PingResponse, error) {
	response := model.PingResponse{}

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
