/**
 * @Author: xianxiong
 * @Date: 2020/10/30 15:35
 */

package model

import (
	"time"
)

// 用户表
type EpUser struct {
	Id         string    `gorm:"column:id;primary_key" json:"id"`            // 主键
	UserName   string    `gorm:"column:user_name;NOT NULL" json:"user_name"` // 用户名
	Password   string    `gorm:"column:password;NOT NULL" json:"password"`   // 密码
	NickName   string    `gorm:"column:nick_name" json:"nick_name"`          // 昵称
	CardId     string    `gorm:"column:card_id" json:"card_id"`              // 身份证id
	Source     string    `gorm:"column:source" json:"source"`                // 用户来源：系统用户、微信用户、天府银行用户
	DataStatus string    `gorm:"column:data_status" json:"data_status"`      // 数据状态（0：有效，1：无效）
	CreateId   int64     `gorm:"column:create_id" json:"create_id"`          // 创建人id
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`      // 创建时间
}

func (m *EpUser) TableName() string {
	return "ep_user"
}
