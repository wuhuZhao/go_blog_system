package blog_handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	blog_strcut "hkzhao/go_blog_system/blog/strcut"
	"hkzhao/go_blog_system/common"
	"net/http"
	"strconv"
	"time"
)

func BlogAddHandler(ctx *gin.Context) {
	title := ctx.PostForm("title")
	topicGroup := ctx.PostForm("topic_group")
	content := ctx.PostForm("content")
	author := ctx.PostForm("author")
	createTime := time.Now()
	updateTime := createTime
	res := blog_strcut.Blog{
		Title:      title,
		TopicGroup: topicGroup,
		CreateTime: createTime,
		UpdateTime: updateTime,
		Content:    content,
		Author:     author,
	}
	result := common.Db.Select("Title", "TopicGroup", "CreateTime", "UpdateTime", "Content", "Author").Create(&res)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	jsonData, err1 := json.Marshal(res)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": string(jsonData),
	})
}
func BlogRemoveHandler(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.PostForm("id"), 10, 32)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	blog := blog_strcut.Blog{}
	result := common.Db.Delete(&blog, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	if result.RowsAffected == 1 {
		jsonData, err2 := json.Marshal(blog)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "error",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": string(jsonData),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
}
func BlogUpdateHandler(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.PostForm("id"), 10, 32)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	blog := blog_strcut.Blog{}
	result := common.Db.First(&blog, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	title := ctx.PostForm("title")
	topicGroup := ctx.PostForm("topic_group")
	content := ctx.PostForm("content")
	author := ctx.PostForm("author")
	blog.Title = title
	blog.TopicGroup = topicGroup
	blog.Content = content
	blog.Author = author
	blog.UpdateTime = time.Now()
	result = common.Db.Save(&blog)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	jsonData, err2 := json.Marshal(blog)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": string(jsonData),
	})
}
func BlogGetHandler(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Query("id"), 10, 32)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	blog := blog_strcut.Blog{}
	result := common.Db.First(&blog, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	jsonData, err2 := json.Marshal(blog)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": string(jsonData),
	})

}
func BlogGetAllHandler(ctx *gin.Context) {
	var blogs []blog_strcut.Blog
	result := common.Db.Find(&blogs)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	jsonData, err := json.Marshal(blogs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": string(jsonData),
	})
}
