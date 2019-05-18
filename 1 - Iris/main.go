package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	app.Get("/ping2", func(ctx iris.Context) {
		ctx.WriteString("Testing")
	})

	app.Get("/ping3", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "hello world",
			"error":   "null",
		})
	})

	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8080"))
}
