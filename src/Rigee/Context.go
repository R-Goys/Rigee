package Rigee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	StatusCode int
	Path       string
	Method     string
	Params     map[string]string //路径上的参数
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Request.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

func (ctx *Context) SetCookie(cookie *http.Cookie) {
	ctx.Writer.Header().Add("Set-Cookie", cookie.String())
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Writer.WriteHeader(code)
	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Writer.WriteHeader(code)
	ctx.Writer.Write([]byte(html))
}
