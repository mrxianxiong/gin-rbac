/**
 * @Author: xianxiong
 * @Date: 2020/11/1 15:56
 */

package router

import (
	"gin-rbac/controller"
	"gin-rbac/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 认证相关的接口
	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
	}
	// 用户相关的接口
	epUser := r.Group("/epuser")
	{
		epUser.POST("/addepuser", controller.AddEpUser)
		epUser.GET("/getuserinfo", middleware.AuthMiddleware(), controller.GetUserInfo)
	}
	// 权限相关的接口
	epRole := r.Group("/eprole")
	{
		epRole.POST("/addeprole", controller.AddEpRole)
	}
	return r
}
