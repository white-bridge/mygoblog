package routers

import (
	"github.com/gin-gonic/gin"
	"asablog/api/bak"
	"asablog/api/front"
)

func ApiRegister( router *gin.Engine){


	//用户登录
	router.POST("/api/front/register",front.Register)
	router.POST("/api/front/login",front.Login)

	//后台
	ba := router.Group("/api/admin")
	{
		//登录
		ba.POST("/login",bak.Verify)
		article := new(bak.Article)
		ba.POST("/articleEdit",article.Edit)

	}

}
