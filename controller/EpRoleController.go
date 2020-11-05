/**
 * @Author: xianxiong
 * @Date: 2020/11/1 17:20
 */

package controller

import (
	"gin-rbac/common/database"
	"gin-rbac/model"
	"github.com/gin-gonic/gin"
)

func AddEpRole(c *gin.Context) {
	db := database.GetDBInstance()

	roleCode := c.PostForm("RoleCode")
	roleName := c.PostForm("RoleName")

	newRole := model.EpRole{
		Id:         123,
		RoleCode:   roleCode,
		RoleName:   roleName,
		DataStatus: "0",
		CreateId:   123,
	}
	db.Create(&newRole)
}
