package controller

import (
	"github.com/kintohub/common-go/utils/testutils"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestCommonGoExampleController_Ping_FailCases(t *testing.T) {
	controller := NewCommonGoExampleController(nil, nil)

	tests := []struct {
		requestJson    string
		errorMsg       string
		panicErrorType int
	}{
		{
			requestJson:    ``,
			errorMsg:       "unexpected end of JSON input",
			panicErrorType: testutils.CLIENT_HTTP_ERROR_TYPE,
		},
		{
			requestJson:    `"bad""json"`,
			errorMsg:       "invalid character '\"' after top-level value",
			panicErrorType: testutils.CLIENT_HTTP_ERROR_TYPE,
		},
		{
			requestJson:    `{"message":""}`,
			errorMsg:       `{"errors":{"error":"message: cannot be blank.","fields":{"message":"cannot be blank"}}}`,
			panicErrorType: testutils.VALIDATION_ERROR_TYPE,
		},
	}

	for _, test := range tests {
		ctx := fasthttp.RequestCtx{}
		ctx.Request.SetBody([]byte(test.requestJson))
		t.Run(test.requestJson, func(t *testing.T) {
			//TODO: I think this can be moved to a utility class for common error testing.
			defer func() {
				testutils.AssertPanicError(t,
					recover(),
					fasthttp.StatusBadRequest,
					test.errorMsg,
					test.panicErrorType)
			}()

			controller.Ping(&ctx)
		})
	}
}

func TestExampleSkipingTest(t *testing.T) {
	t.Skip("This is simply here to show that you should not comment tests, but skip them if you want to disable")
}
