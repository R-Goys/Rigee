package Rigee

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	engine.router.engine = engine
	return engine
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
func (engine *Engine) Run(addr string) (err error) {
	fmt.Println("Starting Rigee Server In " + addr)
	return http.ListenAndServe(addr, engine)
}
