package expresso

import (
	"crypto/tls"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Config holds server configuration settings such as read/write timeouts and maximum header bytes.
type Config struct {
	ReadTimeout    time.Duration // Maximum duration for reading the entire request, including the body.
	WriteTimeout   time.Duration // Maximum duration before timing out writes of the response.
	MaxHeaderBytes int           // Maximum number of bytes the server will read parsing the request header.
}

// App is the main structure of the application, encapsulating the router and server configuration.
type App struct {
	router    *httprouter.Router // HTTP request router.
	Config                       // Server configuration settings.
	TLSConfig *tls.Config        // TLS configuration for HTTPS server.
}

// DefaultApp creates and returns an App instance with default configurations.
// Default Config values:
//   - ReadTimeout: 10 seconds
//   - WriteTimeout: 10 seconds
//   - MaxHeaderBytes: 1MB (1 << 20 bytes)
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

// NewApp creates and returns an App instance with custom configuration settings provided by the user.
func NewApp(c Config, t *tls.Config) App {
	return App{
		router:    httprouter.New(),
		Config:    c,
		TLSConfig: t,
	}
}

// ListenAndServe starts the HTTP server on the specified address with the settings provided in the App's Config.
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

// ListenAndServeTLS starts the HTTPS server with the given certificate and key files on the specified address.
func (a App) ListenAndServeTLS(addr, certFile, keyFile string) error {
	server := http.Server{
		Addr:           addr,
		Handler:        a.router,
		ReadTimeout:    a.Config.ReadTimeout,
		WriteTimeout:   a.Config.WriteTimeout,
		MaxHeaderBytes: a.Config.MaxHeaderBytes,
		TLSConfig:      a.TLSConfig,
	}
	return server.ListenAndServeTLS(certFile, keyFile)
}

// HEAD registers a HEAD request handler for the specified path with optional middleware.
func (a App) HEAD(path string, middlewares ...Middleware) {
	a.router.HEAD(path, handle(middlewares))
}

// OPTIONS registers an OPTIONS request handler for the specified path with optional middleware.
func (a App) OPTIONS(path string, middlewares ...Middleware) {
	a.router.OPTIONS(path, handle(middlewares))
}

// GET registers a GET request handler for the specified path with optional middleware.
func (a App) GET(path string, middlewares ...Middleware) {
	a.router.GET(path, handle(middlewares))
}

// POST registers a POST request handler for the specified path with optional middleware.
func (a App) POST(path string, middlewares ...Middleware) {
	a.router.POST(path, handle(middlewares))
}

// PATCH registers a PATCH request handler for the specified path with optional middleware.
func (a App) PATCH(path string, middlewares ...Middleware) {
	a.router.PATCH(path, handle(middlewares))
}

// PUT registers a PUT request handler for the specified path with optional middleware.
func (a App) PUT(path string, middlewares ...Middleware) {
	a.router.PUT(path, handle(middlewares))
}

// DELETE registers a DELETE request handler for the specified path with optional middleware.
func (a App) DELETE(path string, middlewares ...Middleware) {
	a.router.DELETE(path, handle(middlewares))
}

// ServeStatic serves static files from the provided directory for the specified path.
func (a App) ServeStatic(path string, root http.FileSystem) {
	a.router.ServeFiles(path, root)
}

// HandleNotFound sets up a custom 404 Not Found handler with optional middleware.
func (a App) HandleNotFound(middlewares ...Middleware) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handle(middlewares)(w, r, nil)
	})
	a.router.NotFound = h
}

// HandleError sets up a custom error handler with optional middleware.
func (a App) HandleError(handler func(http.ResponseWriter, *http.Request, interface{})) {
	a.router.PanicHandler = handler
}

// handle is a helper function that processes a list of middleware and invokes them sequentially.
func handle(middlewares []Middleware) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		req := requestFromHttpRequest(r)         // Convert the incoming HTTP request to a custom request type.
		res := responseFromHttpResponseWriter(w) // Convert the response writer to a custom response type.

		if req == nil {
			// Handle errors if the request couldn't be processed.
			res.Status(500).Formatted(r, Formatted{
				Text: &Text{"Error - Unable to process the request"},
				HTML: &HTML{"<html><head><title>Error - Unable to process the request</title></head><body>Error - Unable to process the request</body></html>"},
				JSON: &JSON{
					Data: map[string]interface{}{
						"status": 500,
						"error":  "Unable to process the request",
					},
				},
				XML: &XML{
					Data: struct {
						XMLName xml.Name `xml:"error"`
						Status  int      `xml:"status"`
						Error   string   `xml:"error"`
					}{
						Status: 500,
						Error:  "Unable to process the request",
					},
				},
				YAML: &YAML{
					Data: map[string]interface{}{
						"status": 500,
						"error":  "Unable to process the request",
					},
				},
				Default: &JSON{
					Data: map[string]interface{}{
						"status": 500,
						"error":  "Unable to process the request",
					},
				},
			})

			return
		}

		req.Params = p // Attach the URL parameters to the request.

		// Initialize the context for middleware processing.
		ctx := &Context{
			Request:  req,
			Response: res,
			Extras:   map[interface{}]interface{}{},
			goNext:   false,
			Logger:   NewLogger(req),
		}
		ctx.Response.Context = ctx // Link the response to the context.

		// Execute the middleware chain.
		for _, middleware := range middlewares {
			middleware(ctx)
			if !ctx.goNext {
				break
			}
			ctx.goNext = false
		}

		// Log the response context if needed.
		ctx.Dump()
	}
}
