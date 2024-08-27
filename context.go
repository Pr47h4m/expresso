package expresso

type Context struct {
	*Request
	Response
	Extras map[interface{}]interface{}
	goNext bool
	logger *Logger
}

func (c *Context) Logger() *Logger {
	if c.logger == nil {
		c.logger = NewLogger(c.Request)
	}
	return c.logger
}

func (c *Context) Next() {
	c.goNext = true
}
