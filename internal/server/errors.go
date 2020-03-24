package server

import (
	"common-go-example/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

func writeErrorObject(ctx *fasthttp.RequestCtx, err error, statusCode int) {
	b, _ := json.Marshal(err)
	ctx.Response.SetBody([]byte(fmt.Sprintf(utils.ErrorObjFormat, string(b))))
	ctx.Response.SetStatusCode(statusCode)
}

func writeErrorMessage(ctx *fasthttp.RequestCtx, msg string, statusCode int) {
	ctx.Response.SetBody([]byte(fmt.Sprintf(utils.ErrorMessageFmt, msg)))
	ctx.Response.SetStatusCode(statusCode)
}
