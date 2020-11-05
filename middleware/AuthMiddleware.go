/**
 * @Author: xianxiong
 * @Date: 2020/11/2 11:28
 *  认证中间件
 */

package middleware

import (
	"gin-rbac/common/database"
	"gin-rbac/common/jwt"
	"gin-rbac/model"
	"gin-rbac/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.NoAuthorization(c, "未认证", nil)
			c.Abort()
			return
		}
		// 截取token
		tokenString = tokenString[7:]

		// 验证token
		token, claims, err := jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.NoAuthorization(c, "未认证", nil)
			c.Abort()
			return
		}

		// 验证通过获取claims中的userId
		userId := claims.UserId
		db := database.GetDBInstance()
		var user model.EpUser
		db.First(&user, "id=?", userId)

		// 用户不存在
		if user.Id == "" {
			response.NoAuthorization(c, "未认证", nil)
			c.Abort()
			return
		}
		// 用户存在
		c.Set("user", user)
		c.Next()
	}
}
