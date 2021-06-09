package main

import (
	"github.com/gin-gonic/gin"
	blog_handler "hkzhao/go_blog_system/blog/handler"
	"hkzhao/go_blog_system/common"
	user_handler "hkzhao/go_blog_system/user/handler"
	"log"
	"net/http"
)

func main() {
	err := common.InitMySQL()
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	r.Use(Cors())
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
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
