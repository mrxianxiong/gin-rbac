/**
 * @Author: xianxiong
 * @Date: 2020/10/30 15:52
 */

package controller

import (
	"gin-rbac/common"
	"gin-rbac/dto"
	"gin-rbac/model"
	"gin-rbac/response"
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
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
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
	response.Success(c, nil, "注册成功！")
}

// 获取token中设置的用户
func GetUserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"data": dto.ToEpUserDto(user.(model.EpUser))}, "获取成功！")
}
