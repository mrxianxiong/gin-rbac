/**
 * @Author: xianxiong
 * @Date: 2020/10/30 15:52
 */

package controller

import (
	"gin-rbac/common/database"
	"gin-rbac/dto"
	"gin-rbac/model"
	"gin-rbac/response"
	"gin-rbac/util/md5"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

// 新增用户
func AddEpUser(c *gin.Context) {
	DB := database.GetDBInstance()
	// 获取参数
	userName := c.PostForm("UserName")
	password := c.PostForm("Password")
	nickName := c.PostForm("NickName")
	cardId := c.PostForm("CardId")
	source := c.PostForm("Source")
	createId := c.PostForm("CreateId")

	// 创建用户
	encodeMD5 := md5.EncodeMD5(password)
	create_id, err := strconv.ParseInt(createId, 10, 64)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "转换异常！")
		return
	}
	uuid := guuid.New().String()
	newUser := model.EpUser{
		Id:         uuid,
		UserName:   userName,
		Password:   string(encodeMD5),
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
