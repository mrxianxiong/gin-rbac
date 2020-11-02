/**
 * @Author: xianxiong
 * @Date: 2020/11/2 14:41
 */

package controller

import (
	"gin-rbac/common"
	"gin-rbac/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// 登录
func Login(c *gin.Context) {
	db := common.GetDBInstance()
	// 获取参数
	userName := c.PostForm("userName")
	password := c.PostForm("passWord")

	// 参数验证
	if len(userName) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号必须是11位"})
		return
	}
	// 判断username是否存在
	var user model.EpUser
	db.Where("user_name = ?", userName).First(&user)
	if user.Id == "" || len(user.Id) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "账号不存在"})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		log.Printf("token generate error: %v", err)
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功！",
	})
}
