package main

import (
	"github.com/gin-gonic/gin"
	blog_handler "hkzhao/go_blog_system/blog/handler"
	"hkzhao/go_blog_system/common"
	user_handler "hkzhao/go_blog_system/user/handler"

	"net/http"
)

func main() {
	err := common.InitMySQL()
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	r.Use(CorsHandler())
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
func CorsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "false")
		context.Set("content-type", "application/json") //设置返回格式是json

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, gin.H{
				"msg": "option",
			})
		}

		//处理请求
		context.Next()
	}
}
