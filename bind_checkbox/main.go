package main

import "github.com/gin-gonic/gin"

type FormData struct {
	Hobbies []string `form:"hobbies" binding:"required"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./static/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	r.POST("/submit", func(ctx *gin.Context) {
		var form FormData
		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(500, gin.H{
				"msg": err.Error(),
			})
		}
		ctx.JSON(200, form)
	})

	r.Run(":8080")
}
