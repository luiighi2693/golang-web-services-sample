package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"golang-web-services-sample/comment/commentController"
)

func NewApp() *iris.Application {
	app := iris.New()

	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	commentController.SetEndpoints(app)

	return app
}

func main() {
	app := NewApp()
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}