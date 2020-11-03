/**
 * @Author: xianxiong
 * @Date: 2020/11/3 9:43
 */

package dto

import "gin-rbac/model"

type EpUserDto struct {
	Id       string `gorm:"column:id;primary_key"`     // 主键
	UserName string `gorm:"column:user_name;NOT NULL"` // 用户名
	NickName string `gorm:"column:nick_name"`          // 昵称
	CardId   string `gorm:"column:card_id"`            // 身份证id
	Source   string `gorm:"column:source"`             // 用户来源：系统用户、微信用户、天府银行用户
}

func ToEpUserDto(user model.EpUser) EpUserDto {
	return EpUserDto{
		Id:       user.Id,
		UserName: user.UserName,
		NickName: user.NickName,
		CardId:   user.CardId,
		Source:   user.Source,
	}

}
