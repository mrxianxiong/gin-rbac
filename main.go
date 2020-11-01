/**
 * @Author: xianxiong
 * @Date: 2020/10/30 15:20
 */

package main

import (
	"gin-rbac/common"
	"gin-rbac/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := common.InitDB()
	// 延迟关闭数据库连接
	defer db.Close()
	//
	r := gin.Default()
	router.CollectRoute(r)
	panic(r.Run())

}
