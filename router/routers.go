/**
 * @Author: xianxiong
 * @Date: 2020/11/1 15:56
 */

package router

import (
	"gin-rbac/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	epUser := r.Group("/epuser")
	{
		epUser.POST("/addepuser", controller.AddEpUser)
	}
	epRole := r.Group("/eprole")
	{
		epRole.POST("/addeprole")
	}
	return r
}