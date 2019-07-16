package router

import (
	"common-go-example/internal/controller"
	"github.com/buaazp/fasthttprouter"
)

func New() *fasthttprouter.Router {
	c := controller.NewCommonGoExampleController(nil, nil)
	router := fasthttprouter.New()

	router.Handle("POST", "/ping", c.Ping)

	return router
}
