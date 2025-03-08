package main

import (
	"fmt"
	Rigee "github.com/R-Goys/Rigee/core"
)

func main() {
	r := Rigee.New()
	user := r.Group("/:uid")
	user.Use(func(c *Rigee.Context) {
		fmt.Println("中间件执行")
		c.Next()
		fmt.Println("中间件的返回")
	})
	user.POST("/page", func(c *Rigee.Context) {
		c.JSON(200, Rigee.H{
			"hello":     "ClassOne",
			"status":    200,
			"mymessage": c.PostForm("msg"),
			"param":     c.Params,
		})
	})
	r.Run(":8080")
}
