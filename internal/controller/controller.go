package controller

import (
	"github.com/kintohub/common-go/client"
	"github.com/kintohub/common-go/kintohub"
	"github.com/kintohub/common-go/kintohub/graphql/hasura"
	"github.com/kintohub/common-go/utils/fasthttputils"
	"github.com/kintohub/common-go/utils/json"
	"github.com/valyala/fasthttp"
)

type CommonGoExampleController struct {
	httpClient client.IHttpClient
	graphqlApi kintohub.IGraphqlApi
}

func NewCommonGoExampleController(httpClient client.IHttpClient,
	graphqlApi kintohub.IGraphqlApi) *CommonGoExampleController {

	if httpClient == nil {
		httpClient = client.NewHttpClientCaller(nil)
	}

	if graphqlApi == nil {
		graphqlApi = kintohub.NewClient(
			hasura.NewClient(httpClient),
		)
	}

	return &CommonGoExampleController{
		httpClient: httpClient,
		graphqlApi: graphqlApi,
	}
}

func (c *CommonGoExampleController) Ping(ctx *fasthttp.RequestCtx) {
	request := PingRequest{}
	json.PanicValidateClientBytesToStruct(ctx.PostBody(), &request)
	response := pong(request)
	fasthttputils.WriteJsonResponse(ctx, fasthttp.StatusOK, response)
}
