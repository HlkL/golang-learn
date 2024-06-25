package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./static/*")

	r.GET("/events", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		for {
			c.Writer.Write([]byte(fmt.Sprintf("data: %s\n\n", time.Now().Format("15:04:05"))))
			c.Writer.Flush()
			time.Sleep(2 * time.Second)
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.Run(":8080")
}
