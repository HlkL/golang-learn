package main

import (
	"fmt"
	"net/http"

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

	// 匹配任意请求方式
	r.Any("/any", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "any request method",
		})
	})

	// 定义没有处理函数的路由
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"msg": "page not found",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "This is /users with GET")
	})

	// 定义 r.NoMethod() 在所有路由之后
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"msg": "method not allowed",
		})
	})

	g := r.Group("/g")
	{
		g.GET("/get", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"msg": ctx.Request.URL.Path,
			})
		})

		// 嵌套路由组
		xx := g.Group("/xx")
		xx.GET("oo", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"msg": ctx.Request.URL.Path,
			})
		})

	}

	r.Run()
}
