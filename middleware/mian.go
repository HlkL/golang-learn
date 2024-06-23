package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 使用了默认的中间件 gin.Logger()、gin.Recovery()， 不使用任何中间件可使用 gin.New()
	r := gin.Default()

	// 使用全局中间件
	// r.Use(myMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ok"})
	})

	// 单独使用中间件
	r.GET("/api", myMiddleware(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ok"})
	})

	// 路由组
	g := r.Group("/g")
	g.Use(myMiddleware())
	{
		g.GET("/api", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "ok"})
		})
	}

	// basic auth
	auth := r.Group("basic-auth", gin.BasicAuth(gin.Accounts{
		"foo":    "foo",
		"austin": "123423",
		"lena":   "lena",
	}))

	auth.GET("/secrets", func(ctx *gin.Context) {
		user := ctx.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			ctx.JSON(200, gin.H{"user": user, "secret": secret})
		} else {
			ctx.JSON(403, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	r.Run(":8080")
}

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func myMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		fmt.Println("myMiddleware中间件开始执行")

		// 继续执行后续中间件
		c.Next()
		// 不再执行当前请求后续代码流程
		// c.Abort()

		latencyTime := time.Since(startTime)
		// 记录请求信息
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		fmt.Printf("[myMiddleware] %s %3d %s %13v\n", reqMethod, statusCode, reqURI, latencyTime)
		fmt.Println("myMiddleware执行结束")
	}
}
