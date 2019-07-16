package main

import (
	"common-go-example/internal/router"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kintohub/common-go/logger"
	"github.com/kintohub/common-go/server"
)

func main() {
	log := logger.NewReflectLogger()
	logger.SetLogger(log)
	r := router.New()
	server.NewServer(r).Start()
}
