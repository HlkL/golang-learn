package main

import "github.com/gin-gonic/gin"


func main()  {
	r := gin.Default()

	// 使用 JSONP 向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中。
	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"message": "Hello from the server!",
		}

		// /jsonp?callback=x
		// 将输出：x({\"message\":\"Hello from the server!\"})
		c.JSONP(200, data)
	})

	r.Run(":8080")
}