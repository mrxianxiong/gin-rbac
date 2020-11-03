/**
 * @Author: xianxiong
 * @Date: 2020/11/1 14:29
 */

package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

// 初始化数据库实例
func InitDB() *gorm.DB {
	driverName := viper.GetString("databases.driverName")
	host := viper.GetString("databases.host")
	port := viper.GetString("databases.port")
	database := viper.GetString("databases.database")
	username := viper.GetString("databases.username")
	password := viper.GetString("databases.password")
	charset := viper.GetString("databases.charset")
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
