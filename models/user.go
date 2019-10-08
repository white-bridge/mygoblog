package models

import (
	"crypto/md5"
	"asablog/helper"
	"encoding/hex"
)

type User struct {
	ID             int      `gorm:"PRIMARY_KEY"`
	Nickname       string	`json:"nickname"`
	Email          string	`json:"email"`
	Avatar         string
	Password       string	`json:"password"`
	Salt           string
	Active         int
	Created_at     string
	Updated_at     string
	Isdel          int
}

func (User) TableName() string{
	return "t_blog_user"
}

func (a *User) First(str1 string,str2 string) *User {
	orm.Table("t_blog_user").Where(str1,str2).First(a)
	return a
}

func (_ *User) List() []User{
	var user []User
	orm.Table("t_blog_user").Select("id,nickname,email,avatar,password,salt,active,created_at,updated_at,isdel").Order("id desc").Find(&user)
	return  user
}

func (_ *User) Insert(u *User) int {
	u.Salt = helper.RandStringRunes(4)
	h := md5.New()
	h.Write([]byte(u.Password + u.Salt))
	u.Password = hex.EncodeToString(h.Sum(nil))
	orm.Create(u)
	return u.ID
}