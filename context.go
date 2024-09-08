package expresso

// Context represents the context of a request, holding the request and response objects,
// along with additional data and control flags used during middleware processing.
type Context struct {
	*Request                             // Embedded request object containing details about the incoming request.
	Response                             // Embedded response object for sending data back to the client.
	Extras   map[interface{}]interface{} // A map for storing additional data that may be used across middlewares.
	goNext   bool                        // A flag to control the flow of middleware execution.
	*Logger                              // Logger for logging messages.
}

// Next sets the goNext flag to true, allowing the next middleware in the chain to be executed.
func (c *Context) Next() {
	c.goNext = true
}

func (c *Context) Abort() {
	c.goNext = false
}
