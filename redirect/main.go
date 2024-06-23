package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 外部链接重定向
	r.GET("/r", func(ctx *gin.Context) {
		ctx.Redirect(300, "https://hougen.fun")
	})

	// 内部重定向
	r.GET("/r2", func(ctx *gin.Context) {
		fmt.Println("内部重定向 -> r3")
		ctx.Request.URL.Path = "/r3"
		r.HandleContext(ctx)
	})

	r.GET("/r3", func(ctx *gin.Context) {
		fmt.Println("r3")
		ctx.JSON(200, gin.H{
			"msg": "ok",
		})
	})

	r.Run()
}
