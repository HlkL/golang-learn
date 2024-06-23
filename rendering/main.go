package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/json", func(ctx *gin.Context) {
		user := User{}
		user.Name = "zhangsan"
		user.Age = 13
		ctx.JSON(200, user)
	})

	r.GET("/xml", func(ctx *gin.Context) {
		user := User{}
		user.Name = "zhangsan"
		user.Age = 13
		ctx.XML(200, user)
	})

	r.GET("/yaml", func(ctx *gin.Context) {
		user := User{}
		user.Name = "zhangsan"
		user.Age = 13
		ctx.YAML(200, user)
	})

	// html
	r.LoadHTMLGlob("./templates/*")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"title": "hello,world",
		})

	})

	// file
	r.GET("/images",func(ctx *gin.Context) {
		ctx.File("./static/img.png")
	})

	r.Run()
}
