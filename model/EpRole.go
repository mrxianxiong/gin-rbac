/**
 * @Author: xianxiong
 * @Date: 2020/11/1 17:21
 */

package model

import (
	"time"
)

// 角色表
type EpRole struct {
	Id         int64     `gorm:"column:id;primary_key"`       // 主键
	RoleCode   string    `gorm:"column:role_code;NOT NULL"`   // 角色编码
	RoleName   string    `gorm:"column:role_name;NOT NULL"`   // 角色名称
	DataStatus string    `gorm:"column:data_status;NOT NULL"` // 数据状态（0：有效，1：无效）
	CreateId   int64     `gorm:"column:create_id"`            // 创建人id
	CreateTime time.Time `gorm:"column:create_time"`          // 创建时间
}

func (m *EpRole) TableName() string {
	return "ep_role"
}
