/**
 * @Author: xianxiong
 * @Date: 2020/10/30 15:20
 */

package main

import (
	"gin-rbac/common/database"
	"gin-rbac/common/redis"
	"gin-rbac/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// 加载配置文件
	InitConfig()
	db := database.InitDB()
	// 延迟关闭数据库连接
	defer db.Close()
	// 初始化redis
	redis.InitRedis()
	// 加载路由
	r := gin.Default()
	router.CollectRoute(r)
	// 获取端口信息
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

// 初始化配置文件
func InitConfig() {
	dir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
