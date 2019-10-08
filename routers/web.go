package routers

import (
	"asablog/controllers/front"
	"github.com/gin-gonic/gin"
	"asablog/controllers/bak"
	"asablog/middleware"
	"html/template"
	"time"
)

//时间戳转换成日期
func formatAsDate(t int)  string{
	return time.Unix(int64(t),0).Format("2006-01-02 15:04:05")
}

//类型转换
func formatAsType(t int) string{
	switch t {
		case 1:
			return "PHP"
		case 2:
			return "Go"
		case 3:
			return "Linux"
		case 4:
			return "Js"
		default:
			return "其他"
	}
}

//审核状态转换
func formatAsCheck(t int) string{
	switch t {
	case 0:
		return "编辑中"
	case 1:
		return "待审核"
	case 2:
		return "审核通过"
	case 3:
		return "审核失败"
	default:
		panic("审核状态错误")
	}
}

//字符串转html
func unescaped(x string) interface{} { return template.HTML(x)}

//处理web路由注册
func WebRegister(router *gin.Engine){

	//首页
	router.GET("/",front.Index)
	router.GET("/about",front.About)

	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"formatAsType": formatAsType,
		"formatAsCheck": formatAsCheck,
		"unescaped": unescaped,
	})


	a := router.Group("/articles")
	{
		//a.Use(middleware.TokenVerify())
		articles := new(front.Articles)
		a.GET("/index",articles.Index)
		a.GET("/single",articles.Single)
	}

	//后台
	ba := router.Group("/admin")
	{
		ba.GET("/login",bak.Login)
		ba.Use(middleware.ALogin())
		ba.GET("/",bak.Index)
		articles := new(bak.Articles)
		ba.GET("/articles",articles.Index)
		ba.GET("/articleEdit",articles.Edit)
		user := new(bak.User)
		ba.GET("/users",user.Index)

	}
}
