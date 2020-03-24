package main

import (
	"common-go-example/internal/config"
	"common-go-example/internal/server"
	"github.com/kintohub/utils-go/logger"
)

func main() {
	logger.SetLogLevel(config.LogLevel)
	server.Start()
}
