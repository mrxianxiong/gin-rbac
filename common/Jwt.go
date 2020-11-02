/**
 * @Author: xianxiong
 * @Date: 2020/11/1 21:43
 */

package common

import (
	"gin-rbac/model"
	"gin-rbac/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// token 生成
func ReleaseToken(user model.EpUser) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: util.StrToUInt(user.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "oceanlearn.tech",
			IssuedAt:  time.Now().Unix(),
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 验证token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	return token, Claims, err
}
