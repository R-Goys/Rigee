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
	index      int
	handlers   []HandlerFunc //结构为一堆中间件+最终的接口
	engine     *Engine
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
		index:   -1,
	}
}

func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		for ; c.index < len(c.handlers); c.index++ {
			c.handlers[c.index](c)
		}
	} else {
		c.JSON(http.StatusBadRequest, H{
			"status": 40001,
			"error":  "Next HandlerFunc Not Found!",
		})
	}
	return
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	c.Writer.Header().Add("Set-Cookie", cookie.String())
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.WriteHeader(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.JSON(http.StatusInternalServerError, H{
			"status": 500,
			"error":  err.Error(),
		})
		return
	}
}
