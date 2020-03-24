package server

import (
	"common-go-example/internal/config"
	"github.com/buaazp/fasthttprouter"
	"github.com/kintohub/utils-go/logger"
	"github.com/valyala/fasthttp"
)

func Start(port string) {
	r := newRouter()
	logger.Infof("Starting server on port %v", port)
	logger.Errorf("Fasthttp server crashed %v",
		fasthttp.ListenAndServe(":"+config.ServerPort, r.Handler))
}

func newRouter() *fasthttprouter.Router {
	c := newController()
	router := fasthttprouter.New()

	router.Handle("POST", "/ping", c.ping)

	return router
}
