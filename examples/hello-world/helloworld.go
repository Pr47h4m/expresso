package helloworld

import (
	"github.com/pr47h4m/expresso"
)

func App() *expresso.App {
	app := expresso.DefaultApp()

	app.GET("/", func(ctx *expresso.Context) {
		ctx.Send(expresso.HTML{
			Content: "Hello World",
		})
	})

	return &app
}
