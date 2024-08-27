package expresso

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Config struct {
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type App struct {
	router *httprouter.Router
	Config
}

func DefaultApp() App {
	return App{
		router: httprouter.New(),
		Config: Config{
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func NewApp(c Config) App {
	return App{
		router: httprouter.New(),
		Config: c,
	}
}

func (a App) ListenAndServe(addr string) error {
	server := http.Server{
		Addr:           addr,
		Handler:        a.router,
		ReadTimeout:    a.Config.ReadTimeout,
		WriteTimeout:   a.Config.WriteTimeout,
		MaxHeaderBytes: a.Config.MaxHeaderBytes,
	}
	return server.ListenAndServe()
}

func (a App) ListenAndServeTLS(addr, certFile, keyFile string) error {
	server := http.Server{
		Addr:           addr,
		Handler:        a.router,
		ReadTimeout:    a.Config.ReadTimeout,
		WriteTimeout:   a.Config.WriteTimeout,
		MaxHeaderBytes: a.Config.MaxHeaderBytes,
	}
	return server.ListenAndServeTLS(certFile, keyFile)
}

func (a App) HEAD(path string, middlewares ...Middleware) {
	a.router.HEAD(path, handle(middlewares))
}

func (a App) OPTIONS(path string, middlewares ...Middleware) {
	a.router.OPTIONS(path, handle(middlewares))
}

func (a App) GET(path string, middlewares ...Middleware) {
	a.router.GET(path, handle(middlewares))
}

func (a App) POST(path string, middlewares ...Middleware) {
	a.router.POST(path, handle(middlewares))

}

func (a App) PATCH(path string, middlewares ...Middleware) {
	a.router.PATCH(path, handle(middlewares))

}

func (a App) PUT(path string, middlewares ...Middleware) {
	a.router.PUT(path, handle(middlewares))

}

func (a App) DELETE(path string, middlewares ...Middleware) {
	a.router.DELETE(path, handle(middlewares))
}

func (a App) ServeStatic(path string, root http.Dir) {
	a.router.ServeFiles(path, root)
}

func (a App) HandleNotFound(middlewares ...Middleware) {
	// TODO: HandleNotFound
}

func (a App) HandleError(middlewares ...Middleware) {
	// TODO: HandleError
}

func handle(middlewares []Middleware) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		req := requestFromHttpRequest(r)
		res := responseFromHttpResponseWriter(w)
		if req != nil {
			req.Params = p

			ctx := &Context{
				Request:  req,
				Response: res,
				Extras:   map[interface{}]interface{}{},
				goNext:   false,
			}

			for _, middleware := range middlewares {
				middleware(ctx)
				if !ctx.goNext {
					break
				}
				ctx.goNext = false
			}

			ctx.Logger().Dump()
		} else {
			res.Status(500).Formatted(r, Formatted{
				Text: Text{
					Text: "Error - Unable to process the request",
				},
				HTML: HTML{
					HTML: "<html><head><title>Error - Unable to process the request</title></head><body>Error - Unable to process the request</body></html>",
				},
				JSON: JSON{
					"status": 500,
					"error":  "Unable to process the request",
				},
				XML: XML{
					XML: struct {
						XMLName xml.Name `xml:"error"`
						Status  int      `xml:"status"`
						Error   string   `xml:"error"`
					}{
						Status: 500,
						Error:  "Unable to process the request",
					},
				},
				Default: JSON{
					"status": 500,
					"error":  "Unable to process the request",
				},
			})
		}
	}
}
