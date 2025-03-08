package Rigee

import (
	"log"
	"net/http"
	"path"
)

// RouterGroup 路由组主要可以在前缀相同的路由组中添加中间件
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newgroup := &RouterGroup{
		engine: engine,
		prefix: group.prefix + prefix,
	}
	engine.groups = append(engine.groups, newgroup)
	return newgroup
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) addRoute(method string, cmp string, handler HandlerFunc) {
	pattern := group.prefix + cmp
	log.Printf("AddRoute %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	group.addRoute("PUT", pattern, handler)
}
func (group *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	group.addRoute("DELETE", pattern, handler)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}
