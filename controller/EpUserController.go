/**
 * @Author: xianxiong
 * @Date: 2020/10/30 15:52
 */

package controller

import (
	"gin-rbac/common"
	"gin-rbac/model"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

// 新增用户
func AddEpUser(c *gin.Context) {
	DB := common.GetDBInstance()
	// 获取参数
	userName := c.PostForm("UserName")
	password := c.PostForm("Password")
	nickName := c.PostForm("NickName")
	cardId := c.PostForm("CardId")
	source := c.PostForm("Source")
	createId := c.PostForm("CreateId")

	// 创建用户
	haseDPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	create_id, err := strconv.ParseInt(createId, 10, 64)
	uuid := guuid.New().String()
	newUser := model.EpUser{
		Id:         uuid,
		UserName:   userName,
		Password:   string(haseDPassword),
		NickName:   nickName,
		CardId:     cardId,
		Source:     source,
		DataStatus: "0",
		CreateId:   create_id,
		CreateTime: time.Now(),
	}
	DB.Create(&newUser)
	// 返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功！",
	})
}

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
	token := "123"
	// 返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功！",
	})
}
