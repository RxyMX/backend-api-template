package controller

import (
	"github.com/kintohub/common-go/server/middleware"
	"github.com/kintohub/common-go/utils/json"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestCommonGoExampleController_Ping_FailCases(t *testing.T) {
	controller := NewCommonGoExampleController(nil, nil)

	tests := []struct {
		requestJson string
		errorMsg    string
	}{
		{
			requestJson: ``,
			errorMsg:    "unexpected end of JSON input",
		},
		{
			requestJson: `"bad""json"`,
			errorMsg:    "invalid character '\"' after top-level value",
		},
		{
			requestJson: `{"message":""}`,
			errorMsg:    `{"errors":{"error":"message: cannot be blank.","fields":{"message":"cannot be blank"}}}`,
		},
	}

	for _, test := range tests {
		ctx := fasthttp.RequestCtx{}
		ctx.Request.SetBody([]byte(test.requestJson))
		t.Run(test.requestJson, func(t *testing.T) {
			//TODO: I think this can be moved to a utility class for common error testing.
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code failed to panic on test %+v", test)
				} else if httpError, ok := r.(middleware.ClientHttpError); ok {
					assert.Equal(t, fasthttp.StatusBadRequest, httpError.StatuCode)
					assert.Equal(t, test.errorMsg, httpError.Message)
				} else if validationError, ok := r.(middleware.ValidationError); ok {
					assert.Equal(t, fasthttp.StatusBadRequest, validationError.StatuCode)
					assert.Equal(t, test.errorMsg, json.JsonStructToJsonString(validationError.Body()))
				} else {
					t.Error("Panic is not a client HttpClientError or ValidationError!")
				}
			}()

			controller.Ping(&ctx)
		})
	}
}
