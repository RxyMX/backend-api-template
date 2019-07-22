package controller

import (
	"common-go-example/internal/config"
	"github.com/kintohub/common-go/utils/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func Test_Pong(t *testing.T) {
	tests := []struct {
		Name                string
		request             PingRequest
		response            interface{}
		pongEnabled         bool
		pongOverrideMessage string
	}{
		{
			Name:        "Pong disabled test",
			pongEnabled: false,
			request: PingRequest{
				Message: "hi",
			},
			response: `pong is currently on vacation and cannot be found`,
		},
		{
			Name:        "Ping/Pong successful test",
			pongEnabled: true,
			request: PingRequest{
				Message: "hi",
			},
			response: &PingResponse{
				Message: "hi",
			},
		},
		{
			Name:                "Ping/Pong override message test",
			pongEnabled:         true,
			pongOverrideMessage: "coop was here",
			request: PingRequest{
				Message: "hi",
			},
			response: &PingResponse{
				Message: "coop was here",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			config.PongEnabled = test.pongEnabled
			config.PongOverrideMessage = test.pongOverrideMessage

			if test.pongEnabled == false {
				defer func() {
					testutils.AssertPanicError(t,
						recover(),
						fasthttp.StatusNotFound,
						test.response.(string),
						testutils.CLIENT_HTTP_ERROR_TYPE)
				}()
			}

			response := pong(test.request)
			assert.Equal(t, test.response, response)
		})
	}
}
