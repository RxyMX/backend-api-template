package server

import (
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

func TestExampleSkippingTest(t *testing.T) {
	t.Skip("This is simply here to show that you should not comment tests, but skip them if you want to disable")
}
