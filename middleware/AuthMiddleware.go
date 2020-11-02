/**
 * @Author: xianxiong
 * @Date: 2020/11/2 11:28
 *  认证中间件
 */

package middleware

import (
	"gin-rbac/common"
	"gin-rbac/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "未认证"})
			c.Abort()
			return
		}
		// 截取token
		tokenString = tokenString[7:]

		// 验证token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "未认证"})
			c.Abort()
			return
		}

		// 验证通过获取claims中的userId
		userId := string(claims.UserId)
		db := common.GetDBInstance()
		var user model.EpUser
		db.First(&user, userId)

		// 用户不存在
		if user.Id == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "未认证"})
			c.Abort()
			return
		}
		// 用户存在
		c.Set("user", user)
		c.Next()
	}
}
