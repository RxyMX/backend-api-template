package controller

import (
	"github.com/kintohub/common-go/client"
	"github.com/kintohub/common-go/kintohub"
	"github.com/kintohub/common-go/kintohub/graphql/hasura"
	"github.com/kintohub/common-go/utils/fasthttputils"
	"github.com/kintohub/common-go/utils/json"
	"github.com/valyala/fasthttp"
)

type Controller struct {
	httpClient client.IHttpClient
	graphqlApi kintohub.IGraphqlApi
}

func New(httpClient client.IHttpClient,
	graphqlApi kintohub.IGraphqlApi) *Controller {

	if httpClient == nil {
		httpClient = client.NewHttpClientCaller(nil)
	}

	if graphqlApi == nil {
		graphqlApi = kintohub.NewClient(
			hasura.NewClient(httpClient),
		)
	}

	return &Controller{
		httpClient: httpClient,
		graphqlApi: graphqlApi,
	}
}

func (c *Controller) Ping(ctx *fasthttp.RequestCtx) {
	request := PingRequest{}
	json.PanicValidateClientBytesToStruct(ctx.PostBody(), &request)
	response := c.pong(request)

  fasthttputils.WriteJsonResponse(ctx, fasthttp.StatusOK, response)
}
