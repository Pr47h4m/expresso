package expresso

import "sync"

// Context represents the context of a request, holding the request and response objects,
// along with additional data and control flags used during middleware processing.
type Context struct {
	*Request                             // Embedded request object containing details about the incoming request.
	Response                             // Embedded response object for sending data back to the client.
	Extras   map[interface{}]interface{} // A map for storing additional data that may be used across middlewares.
	goNext   bool                        // A flag to control the flow of middleware execution.
	logger   *Logger                     // A lazy-loaded logger associated with the current request.
	once     sync.Once                   // Ensures the logger is initialized only once.
}

// Logger returns the logger associated with the current context.
// If the logger is not yet initialized, it creates a new one using the request details.
func (c *Context) Logger() *Logger {
	c.once.Do(func() {
		c.logger = NewLogger(c.Request) // Lazy-load the logger if it hasn't been initialized.
	})
	return c.logger
}

// Next sets the goNext flag to true, allowing the next middleware in the chain to be executed.
func (c *Context) Next() {
	c.goNext = true
}
