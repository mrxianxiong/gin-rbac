/**
 * @Author: xianxiong
 * @Date: 2020/11/2 14:41
 */

package controller

import (
	"encoding/json"
	"gin-rbac/common/database"
	"gin-rbac/common/jwt"
	"gin-rbac/common/redis"
	"gin-rbac/model"
	"gin-rbac/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// 登录
func Login(c *gin.Context) {
	db := database.GetDBInstance()
	// 获取参数
	userName := c.PostForm("userName")
	password := c.PostForm("passWord")

	// 参数验证
	if len(userName) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "账号必须是11位!")
		return
	}
	// 判断username是否存在
	var user model.EpUser
	db.Where("user_name = ?", userName).First(&user)
	if user.Id == "" || len(user.Id) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "账号不存在!")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 500, nil, "密码错误!")
		return
	}
	// 发放token
	token, err := jwt.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误!")
		log.Printf("token generate error: %v", err)
		return
	}
	// 存储token,user到redis
	userJson, _ := json.Marshal(user)
	b := redis.Set("go-gin:"+token, userJson)
	if !b {
		response.Fail(c, "token放置失败", gin.H{"data": nil})
	}

	// 返回结果
	response.Success(c, gin.H{"data": token}, "登录成功!")
}
