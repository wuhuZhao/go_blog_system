package main

import (
	"github.com/gin-gonic/gin"
	blog_handler "hkzhao/go_blog_system/blog/handler"
	blog_strcut "hkzhao/go_blog_system/blog/strcut"
	"hkzhao/go_blog_system/common"
	user_handler "hkzhao/go_blog_system/user/handler"
)

func main() {
	err := common.InitMySQL()
	if err != nil {
		panic("failed to connect database")
	}
	common.Db.AutoMigrate(&blog_strcut.Blog{})
	common.Db.AutoMigrate(&blog_strcut.TopicGroup{})
	r := gin.Default()
	blogRestApi := r.Group("/blog")
	{
		blogRestApi.GET("/getOne", blog_handler.BlogGetHandler)
		blogRestApi.GET("/getAll", blog_handler.BlogGetAllHandler)
		blogRestApi.POST("/update", blog_handler.BlogUpdateHandler)
		blogRestApi.POST("/remove", blog_handler.BlogRemoveHandler)
		blogRestApi.POST("/add", blog_handler.BlogAddHandler)
	}
	topicGroupRestApi := r.Group("/topic")
	{
		topicGroupRestApi.GET("/getAll", blog_handler.TopicGroupGetAllHandler)
		topicGroupRestApi.POST("/update", blog_handler.TopicGroupUpdateHandler)
		topicGroupRestApi.POST("/add", blog_handler.TopicGroupAddHandler)
		topicGroupRestApi.POST("/remove", blog_handler.TopicGroupRemoveHandler)
	}
	userRestApi := r.Group("/user")
	{
		userRestApi.POST("/login", user_handler.LoginHandler)
	}
	r.Run(":80")
}
