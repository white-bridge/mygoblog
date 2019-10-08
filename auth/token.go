package auth

import (
	"asablog/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 生成前台用户token
func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"nickname": user.Nickname,
		"email": user.Email,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
	})

	return token.SignedString([]byte("frontsecret"))//对应的字符串请自行生成，最后足够使用加密后的字符串
}

// 生成后台登录token
func GenerateBakToken(admin *models.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": admin.ID,
		"name": admin.Name,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
	})

	return token.SignedString([]byte("baksecret"))//对应的字符串请自行生成，最后足够使用加密后的字符串
}

