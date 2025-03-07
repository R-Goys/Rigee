package Rigee

type HandlerFunc func(c *Context)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	r.handlers[method+"-"+pattern] = handler
}

func (r *router) handle(c *Context) {
	key := c.Request.Method + "-" + c.Request.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
		return
	}
	c.JSON(404, H{
		"message": "Not found",
		"status":  404,
	})
	return
}
