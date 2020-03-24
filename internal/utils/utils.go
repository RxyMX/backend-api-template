package utils

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/valyala/fasthttp"
)

const (
	ErrorObjFormat  = `{"errors":%v}`
	ErrorMessageFmt = `{"errors":{"error":"%s"}}`
)

func UnmarshalAndValidate(data []byte, v validation.Validatable) error {
	err := json.Unmarshal(data, v)

	if err == nil {
		err = v.Validate()
	}

	return err
}

func WriteErrorObject(ctx *fasthttp.RequestCtx, err error, statusCode int) {
	b, _ := json.Marshal(err)
	ctx.Response.SetBody([]byte(fmt.Sprintf(ErrorObjFormat, string(b))))
	ctx.Response.SetStatusCode(statusCode)
}

func WriteErrorMessage(ctx *fasthttp.RequestCtx, msg string, statusCode int) {
	ctx.Response.SetBody([]byte(fmt.Sprintf(ErrorMessageFmt, msg)))
	ctx.Response.SetStatusCode(statusCode)
}
