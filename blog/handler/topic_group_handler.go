package blog_handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	blog_strcut "hkzhao/go_blog_system/blog/strcut"
	"hkzhao/go_blog_system/common"
	"net/http"
	"strconv"
)

func TopicGroupAddHandler(c *gin.Context) {
	topic := c.PostForm("topic_group")
	topicGroup := blog_strcut.TopicGroup{
		TOPIC: topic,
	}
	result := common.Db.Create(&topic)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	jsonData, err := json.Marshal(topicGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": string(jsonData),
	})
}

func TopicGroupRemoveHandler(c *gin.Context) {
	id, err1 := strconv.ParseInt(c.PostForm("id"), 10, 32)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	topicGroup := blog_strcut.TopicGroup{}
	result := common.Db.Delete(&topicGroup, id)
	go func(tg *blog_strcut.TopicGroup) {
		common.Db.Where("topic_group=?", tg.TOPIC).Delete(blog_strcut.TopicGroup{})
	}(&topicGroup)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	if result.RowsAffected == 1 {
		jsonData, err2 := json.Marshal(topicGroup)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "error",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": string(jsonData),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
}

func TopicGroupUpdateHandler(c *gin.Context) {
	id, err1 := strconv.ParseInt(c.PostForm("id"), 10, 32)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	newTopic := c.PostForm("topic_group")
	topicGroup := blog_strcut.TopicGroup{}
	result := common.Db.First(&topicGroup, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	go func(tg *blog_strcut.TopicGroup, newVal string) {
		common.Db.Model(blog_strcut.Blog{}).Where("topic_group=?", tg.TOPIC).Updates(blog_strcut.Blog{TopicGroup: newVal})
	}(&topicGroup, newTopic)
	topicGroup.TOPIC = newTopic
	result = common.Db.Save(&topicGroup)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	if result.RowsAffected == 1 {
		jsonData, err2 := json.Marshal(topicGroup)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "error",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": string(jsonData),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
}

func TopicGroupGetAllHandler(c *gin.Context) {
	var topicGroups []blog_strcut.TopicGroup
	result := common.Db.Find(&topicGroups)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	jsonData, err := json.Marshal(topicGroups)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": string(jsonData),
	})
}
