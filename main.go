package main

import (
	"github.com/Rinai-R/Rigee/src/Rigee"
)

func main() {
	r := Rigee.New()
	r.GET("/hello", func(c *Rigee.Context) {
		c.JSON(200, Rigee.H{
			"hello":  "world",
			"status": 200,
		})
	})
	r.Run(":8080")
}
