package controller

import (
	"common-go-example/internal/config"
	"github.com/kintohub/common-go/server/middleware"
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
			pongEnabled: false,
			request: PingRequest{
				Message: "hi",
			},
			response: `pong is currently on vacation and cannot be found`,
		},
		{
			pongEnabled: true,
			request: PingRequest{
				Message: "hi",
			},
			response: &PingResponse{
				Message: "hi",
			},
		},
		{
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

			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code failed to panic on test %+v", test)
				} else if httpError, ok := r.(middleware.ClientHttpError); ok {
					assert.Equal(t, fasthttp.StatusNotFound, httpError.StatuCode)
					assert.Equal(t, test.response, httpError.Message)
				} else {
					t.Error("Panic is not a client HttpClientError or ValidationError!")
				}
			}()

			response := pong(test.request)
			assert.Equal(t, test.response, response)
		})
	}
}
