package Rigee

import (
	"fmt"
	"github.com/Rinai-R/Rigee/src/Rigee/trie"
	"strings"
	"time"
)

type HandlerFunc func(c *Context)

type router struct {
	roots    map[string]*trie.Node
	handlers map[string]HandlerFunc
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*trie.Node),
	}
}

func parseParts(pattern string) []string {
	vs := strings.Split(pattern, "/")
	var parts []string
	for _, v := range vs {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parseParts(pattern)
	key := method + "-" + pattern
	if _, ok := r.handlers[key]; !ok {
		r.roots[method] = &trie.Node{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	fmt.Println("Key", key)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*trie.Node, map[string]string) {
	SearchParts := parseParts(path)
	params := map[string]string{}
	root, ok := r.roots[method]
	if !ok {
		//不存在对应的方法，直接返回
		return nil, nil
	}
	n := root.Search(SearchParts, 0)

	if n != nil {
		parts := parseParts(n.Pattern)
		for idx, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = SearchParts[idx]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(SearchParts[idx:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	fmt.Println(time.Now(), c.Method, c.Path)
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.Pattern
		c.Params = params
		if handler, ok := r.handlers[key]; ok {
			handler(c)
			return
		}
	}
	c.JSON(404, H{
		"message": "Not found",
		"status":  404,
	})
	return
}
