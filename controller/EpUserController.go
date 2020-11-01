/**
 * @Author: xianxiong
 * @Date: 2020/10/30 15:52
 */

package controller

import (
	"gin-rbac/common"
	"gin-rbac/model"
	"github.com/gin-gonic/gin"
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

	// 创建用户
	newUser := model.EpUser{
		Id:         123456,
		UserName:   userName,
		Password:   password,
		NickName:   nickName,
		CardId:     cardId,
		Source:     source,
		DataStatus: "0",
	}
	DB.Create(&newUser)
}
