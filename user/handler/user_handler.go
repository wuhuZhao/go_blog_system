package user_handler

import (
	"github.com/gin-gonic/gin"
	"hkzhao/go_blog_system/common"
	user2 "hkzhao/go_blog_system/user/struct"
	"net/http"
	"strings"
)

func LoginHandler(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
	if len(username) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
		return
	}
	user := user2.Admin{}
	result := common.Db.Where("username=?", username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error",
		})
		return
	}
	if strings.Compare(user.Password, password) == 0 {
		if token, err := ToJwt(username); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": token,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"msg":  "fail",
				"data": username,
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "fail",
			"data": username,
		})
	}
}
