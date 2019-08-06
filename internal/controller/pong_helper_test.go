package controller

import (
	"common-go-example/internal/config"
	"github.com/kintohub/common-go/server/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func Test_ProcessPong(t *testing.T) {
	tests := []struct {
		Name                string
		Message             string
		ResponseMessage     string
		pongEnabled         bool
		pongOverrideMessage string
	}{
		{
			Name:            "panicProcessPong pong disabled test",
			pongEnabled:     false,
			Message:         "hi",
			ResponseMessage: `pong is currently on vacation and cannot be found`,
		},
		{
			Name:            "panicProcessPong ping/pong successful test",
			pongEnabled:     true,
			Message:         "hi",
			ResponseMessage: "hi",
		},
		{
			Name:                "panicProcessPong ping/pong override message test",
			pongEnabled:         true,
			pongOverrideMessage: "roman was here",
			Message:             "hi",
			ResponseMessage:     "roman was here",
		},
	}

	c := New(nil, nil)

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			config.PongEnabled = test.pongEnabled
			config.PongOverrideMessage = test.pongOverrideMessage

			if test.pongEnabled == false {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code failed to panic on test %+v", test)
					} else if httpError, ok := r.(middleware.ClientHttpError); ok {
						assert.Equal(t, fasthttp.StatusNotFound, httpError.StatuCode)
						assert.Equal(t, test.ResponseMessage, httpError.Message)
					} else {
						t.Error("Panic is not a client HttpClientError or ValidationError!")
					}
				}()
			}

			response := c.panicProcessPong(test.Message)
			assert.Equal(t, test.ResponseMessage, response)
		})
	}
}
