/**
 * @Author: xianxiong
 * @Date: 2020/11/1 14:29
 */

package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 初始化数据库实例
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "gin-gorm-project"
	username := "root"
	password := "root"
	charset := "utf8"
	arges := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, arges)
	db.LogMode(true)
	if err != nil {
		panic("failed to connect databases,err:" + err.Error())
	}
	DB = db
	return db
}

// 定义一个方法获取db实例
func GetDBInstance() *gorm.DB {
	return DB
}
