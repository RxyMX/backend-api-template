package main

import (
	"common-go-example/internal/router"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kintohub/common-go/server"
)

func main() {
	r := router.New()
	server.NewServer(r).Start()
}
