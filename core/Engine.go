package Rigee

import (
	"fmt"
	"html/template"
	"net/http"
)

// Engine 启动http服务器的工具罢了
type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	c.engine = engine
	engine.router.handle(c)
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	engine.router.engine = engine
	return engine
}

func (engine *Engine) Run(addr string) (err error) {
	fmt.Println("Starting Rigee Server In " + addr)
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
