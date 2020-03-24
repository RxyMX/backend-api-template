package main

import (
	"common-go-example/internal/config"
	"common-go-example/internal/router"
	"github.com/kintohub/utils-go/logger"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()
	logger.SetLogLevel(config.LogLevel)
	logger.Infof("Successfully started server listening to port %v", config.ServerPort)
	logger.Errorf("Fasthttp server crashed %v",
		fasthttp.ListenAndServe(":"+config.ServerPort, r.Handler))
}
