package controller

import "github.com/kataras/iris"

func IndexController(app iris.Application) {
	app.Get("/index", func(ctx iris.Context) {

	})
}
