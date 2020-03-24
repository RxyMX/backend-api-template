package server

import (
	"common-go-example/internal/model"
	"common-go-example/internal/store"
	"common-go-example/internal/utils"
	"encoding/json"
	"github.com/kintohub/utils-go/logger"
	"github.com/valyala/fasthttp"
)

type Controller struct {
}

func newController() Controller {
	return Controller{}
}

func (c *Controller) ping(ctx *fasthttp.RequestCtx) {
	request := model.PingRequest{}

	err := utils.UnmarshalAndValidate(ctx.PostBody(), &request)

	if err != nil {
		logger.Debugf("Invalid ping request received %v", string(ctx.Request.Body()))
		writeErrorObject(ctx, err, fasthttp.StatusBadRequest)
		return
	}

	logger.Debugf("Processing ping request %v", request)
	response, err := store.Pong(request)

	if err == nil {
		logger.Debugf("Successful ping response %v", response)
		respData, _ := json.Marshal(response)
		ctx.Response.SetBody(respData)
	} else {
		writeErrorMessage(ctx, err.Error(), fasthttp.StatusNotFound)
	}
}
