package models

type Articles struct {
	ID int
	Uid int
	Type int
	Title string
	Content string
	Count int
	Check_status int
	Publish_time int
	Created_at int
	Updated_at int
	Isdel int
}

func (Articles) TableName() string{
	return "t_blog_article"
}

func (a *Articles) First(id int) *Articles {
	orm.Table("t_blog_article").Where(&Articles{ID:id}).First(a)
	return a
}

func (_ *Articles) List() []Articles{
	var articles []Articles
	orm.Table("t_blog_article").Select("id,uid,type,title,content,count,check_status,publish_time,created_at,updated_at,isdel").Order("id desc").Find(&articles)
	return  articles
}