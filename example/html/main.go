package main

import (
	"fmt"
	Rigee "github.com/R-Goys/Rigee/core"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := Rigee.New()

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *Rigee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *Rigee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", Rigee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *Rigee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", Rigee.H{
			"title": "Rigee",
			"now":   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Nanosecond(), time.UTC),
		})
	})

	r.Run(":8080")
}
