package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	// 路由
	r.PUT("/put", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte{65, 66, 67, 68, 69})
	})

	// 路径参数 json输出
	r.POST("user/:name/:age", func(ctx *gin.Context) {
		name := ctx.Param("name")
		age := ctx.Param("age")

		rst := map[string]interface{}{
			"name": name,
			"age":  age,
		}

		ctx.AsciiJSON(http.StatusOK, rst)
	})

	r.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/xml", func(ctx *gin.Context) {
		ctx.XML(200, gin.H{
			"message": "ok",
		})
	})

	// url参数
	r.GET("/user", func(ctx *gin.Context) {
		name := ctx.Query("name")
		ctx.String(http.StatusOK, fmt.Sprintf("name is %s", name))
	})

	// 表单参数
	r.POST("/form", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		age := ctx.PostForm("age")
		idcard := ctx.DefaultPostForm("idcard", "00000001")

		rst := fmt.Sprintf("name is %s, age is %s, idcard is %s", name, age, idcard)
		ctx.Writer.Write([]byte(rst))
	})

	// 解析 application/json 数据
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.String(http.StatusExpectationFailed, err.Error())
		}
		// 使用解析后的 user 数据
		rst := fmt.Sprintf("name is %s, age is %d", user.Name, user.Age)
		c.Writer.Write([]byte(rst))
	})

	//解析 application/xml 数据
	r.POST("/user/xml", func(c *gin.Context) {
		var user XmlUser
		if err := c.BindXML(&user); err != nil {
			c.String(http.StatusExpectationFailed, err.Error())
		}
		// 使用解析后的 user 数据
		rst := fmt.Sprintf("name is %s, age is %d", user.Name, user.Age)
		c.Writer.Write([]byte(rst))
	})

	// 直接解析json/xml数据
	r.POST("/data", func(c *gin.Context) {
		xmlData, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusExpectationFailed, err.Error())
		}
		c.String(http.StatusOK, string(xmlData))
	})

	// 参数动态解析
	r.POST("dynamic", func(ctx *gin.Context) {
		var user DynamicUser
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.String(http.StatusExpectationFailed, err.Error())
		}
		ctx.JSON(200, gin.H{
			"name": user.Name,
			"age":  user.Age,
		})
	})

	// /dynamic?name=&age=
	r.GET("dynamic", func(ctx *gin.Context) {
		var user DynamicUser
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.String(http.StatusExpectationFailed, err.Error())
		}
		ctx.JSON(200, gin.H{
			"name": user.Name,
			"age":  user.Age,
		})
	})

	// 单文件上传
	r.POST("upload", func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusExpectationFailed, gin.H{
				"message": err.Error(),
			})
		}

		fmt.Printf("file.Filename: %v\n", file.Filename)
		// 将上传的文件保存到指定位置
		if err := ctx.SaveUploadedFile(file, "/Users/hougen/Downloads/"+file.Filename); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.POST("upload/multi", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusExpectationFailed, gin.H{
				"message": err.Error(),
			})
		}

		files := form.File["file"]
		for _, file := range files {

			fmt.Printf("file.Filename: %v\n", file.Filename)
			// 将上传的文件保存到指定位置
			if err := ctx.SaveUploadedFile(file, "/Users/hougen/Downloads/"+file.Filename); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run()
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type XmlUser struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

type DynamicUser struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age  int    `form:"age" json:"age" binding:"required"`
}
