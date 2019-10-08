package models

type Admin struct {
	ID         		int
	Name       		string   `json:"username"`
	Password   		string   `json:"password"`
	Salt       		string
	Created_at	    int
	Updated_at      int
	Isdel           int
}

func (Admin) TableName() string{
	return "t_blog_admin"
}

func (a *Admin) First(str1 string,str2 string) *Admin {
	orm.Table("t_blog_admin").Where(str1,str2).First(a)
	return a
}

func (_ *Admin) List() []Admin{
	var admin []Admin
	orm.Table("t_blog_admin").Select("id,name,password,salt,created_at,updated_at,isdel").Order("id desc").Find(&admin)
	return  admin
}
