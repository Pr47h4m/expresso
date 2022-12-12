package expresso

type Context struct {
	*Request
	Response
	Extras map[interface{}]interface{}
	goNext bool
}

func (c *Context) Next() {
	c.goNext = true
}
