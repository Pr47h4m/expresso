package basicrouting

import (
	"github.com/pr47h4m/expresso"
)

func App() *expresso.App {
	app := expresso.DefaultApp()

	app.GET("/", func(ctx *expresso.Context) {
		log := ctx.Logger()
		log.Info("Hello World")
		ctx.Send(expresso.HTML{
			Content: "Hello World",
		})
	})

	app.POST("/", func(ctx *expresso.Context) {
		log := ctx.Logger()
		log.Info("Got a POST request")
		ctx.Send(expresso.HTML{
			Content: "Got a POST request",
		})
	})

	app.PUT("/user", func(ctx *expresso.Context) {
		log := ctx.Logger()
		log.Info("Got a PUT request as /user")
		ctx.Send(expresso.HTML{
			Content: "Got a PUT request as /user",
		})
	})

	app.DELETE("/user", func(ctx *expresso.Context) {
		log := ctx.Logger()
		log.Info("Got a DELETE request at /user")
		ctx.Send(expresso.HTML{
			Content: "Got a DELETE request at /user",
		})
	})

	return &app
}
