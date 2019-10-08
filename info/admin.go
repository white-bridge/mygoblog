package info

type AdminUser struct {
	Id int
	Name string
	Password string
	Salt string
	Isdel int
} 

var AUser  []AdminUser
var State = make(map[string]interface{})