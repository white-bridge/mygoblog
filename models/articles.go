package models

import (
	"time"
)

type Articles struct {
	Id int               `json:"id,string"`
	Uid int              `json:"uid,string"`
	Type int             `json:"type,string"`
	Title string		 `json:"title"`
	Content string		 `json:"content"`
	Count int            `json:"count,string"`
	Check_status int     `json:"check_status,string"`
	Introduce string	 `json:"introduce"`
	Publish_time int
	Created_at string
	Updated_at string
	Isdel int            `json:"isdel,string"`
}

func (Articles) TableName() string{
	return "t_blog_article"
}

func (a *Articles) First(id int) *Articles {
	orm.Table("t_blog_article").Where(&Articles{Id:id}).First(a)
	return a
}

func (_ *Articles) List() []Articles{
	var articles []Articles
	orm.Table("t_blog_article").Order("id desc").Find(&articles)
	return  articles
}

func (_ *Articles) ListAllCondition(arg map[string]interface{}) []Articles{
	orm.LogMode(true)
	var articles []Articles
	if err := orm.Debug().Table("t_blog_article").Select("id,uid,type,title,content,count,check_status,publish_time,created_at,updated_at,isdel,introduce").Where(arg).Order("id desc").Find(&articles).Error; err != nil {
		panic(err)
	}
	return articles
}

func (_ *Articles) ListCondition(limit,offset int,arg map[string]interface{}) []Articles{
	var articles []Articles
	orm.Table("t_blog_article").Select("id,uid,type,title,content,count,check_status,publish_time,created_at,updated_at,isdel,introduce").Where(arg).Limit(limit).Offset(offset).Select("id,uid,type,title,content,count,check_status,publish_time,created_at,updated_at,isdel,introduce").Order("id desc").Find(&articles)
	return  articles
}

// 返回数据插入成功后的ID
func (_ *Articles) Insert(a *Articles) int {
	a.Created_at = time.Now().Format("2006-01-02 15:04:05")
	a.Updated_at = time.Now().Format("2006-01-02 15:04:05")
	orm.Create(a)
	return a.Id
}

// 返回受影响行数
func (article *Articles) Edit(a *Articles) int64 {
	//bakA := a
	//fmt.Println("bakA be :",bakA)
	//ret := article.First(a.Id)
	//// 查无结果 ret为空的Article
	//if ret.Id == 0 {
	//	return 0
	//}
	//fmt.Println("bakA :",bakA)
	rowsAffected := orm.Model(article).Updates(a).RowsAffected
	return rowsAffected
}


//
func (a *Articles) CountRows(arg map[string]interface{}) int {
	total := 0
	orm.Table("t_blog_article").Model(a).Where(arg).Count(&total)
	return total
}