package main

import (
	"github.com/gin-gonic/gin"
	"hkzhao/go_blog_system/common"
)

func main() {
	db, err := common.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
