package main

import (
	"github.com/Rinai-R/Rigee/src/Rigee"
)

func main() {
	r := Rigee.New()
	r.GET("/:hello", func(c *Rigee.Context) {

		c.JSON(200, Rigee.H{
			"hello":     "world",
			"status":    200,
			"mymessage": c.PostForm("msg"),
			"param":     c.Params,
		})
	})
	CLassOne := r.Group("/ClassOne")
	CLassOne.POST("/:hello", func(c *Rigee.Context) {
		c.JSON(200, Rigee.H{
			"hello":     "ClassOne",
			"status":    200,
			"mymessage": c.PostForm("msg"),
			"param":     c.Params,
		})
	})
	r.Run(":8080")
}
