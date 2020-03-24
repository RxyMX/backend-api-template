package server

import (
	"common-go-example/internal/config"
	"common-go-example/internal/model"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestCommonGoExampleController_Ping_FailCases(t *testing.T) {
	controller := newController()
	tests := []struct {
		requestJson string
		errorMsg    string
		statusCode  int
	}{
		{
			requestJson: ``,
			errorMsg:    `{"errors":{"Offset":0}}`,
			statusCode:  fasthttp.StatusBadRequest,
		},
		{
			requestJson: `"bad""json"`,
			errorMsg:    `{"errors":{"Offset":6}}`,
			statusCode:  fasthttp.StatusBadRequest,
		},
		{
			requestJson: `{"message":""}`,
			errorMsg:    `{"errors":{"message":"cannot be blank"}}`,
			statusCode:  fasthttp.StatusBadRequest,
		},
	}

	for _, test := range tests {
		ctx := fasthttp.RequestCtx{}
		ctx.Request.SetBody([]byte(test.requestJson))
		t.Run(test.requestJson, func(t *testing.T) {
			controller.ping(&ctx)

			assert.Equal(t, ctx.Response.StatusCode(), test.statusCode)
			assert.Equal(t, test.errorMsg, string(ctx.Response.Body()))
		})
	}
}

func Test_Pong(t *testing.T) {
	tests := []struct {
		Name                string
		request             model.PingRequest
		response            interface{}
		pongEnabled         bool
		pongOverrideMessage string
		err                 error
	}{
		{
			Name:        "Pong disabled test",
			pongEnabled: false,
			request: model.PingRequest{
				Message: "hi",
			},
			err: errors.New("pong is currently on vacation and cannot be found"),
		},
		{
			Name:        "Ping/Pong successful test",
			pongEnabled: true,
			request: model.PingRequest{
				Message: "hi",
			},
			response: &model.PingResponse{
				Message: "hi",
			},
		},
		{
			Name:                "Ping/Pong override message test",
			pongEnabled:         true,
			pongOverrideMessage: "coop was here",
			request: model.PingRequest{
				Message: "hi",
			},
			response: &model.PingResponse{
				Message: "coop was here",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			config.PongEnabled = test.pongEnabled
			config.PongOverrideMessage = test.pongOverrideMessage

			response, err := pong(test.request)

			assert.Equal(t, test.err, err)
			if err == nil {
				assert.Equal(t, test.response, response)
			}
		})
	}
}

func TestExampleSkippingTest(t *testing.T) {
	t.Skip("This is simply here to show that you should not comment tests, but skip them if you want to disable")
}
