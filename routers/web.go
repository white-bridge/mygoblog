package routers

import (
	"asablog/controllers/front"
	"github.com/gin-gonic/gin"
)

//处理web路由注册
func WebRegister(router *gin.Engine){

	//首页
	router.GET("/",front.Index)

	articles := new(front.Articles)
	a := router.Group("/articles")
	{
		a.GET("/index",articles.Index)
	}
}
