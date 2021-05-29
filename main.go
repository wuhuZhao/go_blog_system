package main

import (
	"github.com/gin-gonic/gin"
	blog_handler "hkzhao/go_blog_system/blog/handler"
	"hkzhao/go_blog_system/common"
)

func main() {
	err := common.InitMySQL()
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/v1/test", blog_handler.BlogAddHandler)
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
