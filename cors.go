package expresso

type Cors struct {
	Path    string
	Origin  string
	Headers string
	Methods string
}

func NewCorsHandler(cors Cors) Middleware {
	return func(ctx *Context) {
		ctx.Response.Headers.Set("access-control-allow-origin", cors.Origin)
		ctx.Response.Headers.Set("access-control-allow-methods", cors.Methods)
		ctx.Response.Headers.Set("access-control-allow-headers", cors.Methods)
		ctx.Next()
	}
}
