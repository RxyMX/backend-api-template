package controller

import (
	"common-go-example/internal/config"
	"common-go-example/internal/utils"
	"encoding/json"
	"errors"
	"github.com/kintohub/utils-go/logger"
	"github.com/valyala/fasthttp"
)

type Controller struct {
}

func New() Controller {
	return Controller{}
}

func (c *Controller) Ping(ctx *fasthttp.RequestCtx) {
	request := PingRequest{}

	err := utils.UnmarshalAndValidate(ctx.PostBody(), &request)

	if err != nil {
		logger.Debugf("Invalid ping request received %v", string(ctx.Request.Body()))
		utils.WriteErrorObject(ctx, err, fasthttp.StatusBadRequest)
		return
	}

	logger.Debugf("Processing ping request %v", request)
	response, err := pong(request)

	if err == nil {
		logger.Debugf("Successful ping response %v", response)
		respData, _ := json.Marshal(response)
		ctx.Response.SetBody(respData)
	} else {
		utils.WriteErrorMessage(ctx, err.Error(), fasthttp.StatusNotFound)
	}
}

func pong(request PingRequest) (*PingResponse, error) {
	response := PingResponse{}

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
