package expresso

// Middleware represents a function that processes an HTTP request within a specific context.
// It takes a pointer to a Context object, allowing the middleware to modify the request, response,
// or other aspects of the context as needed. Middlewares can be chained together to handle requests
// in a modular and reusable way.
type Middleware func(*Context)
