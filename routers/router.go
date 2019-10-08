package routers

import "github.com/gin-gonic/gin"

func Register() *gin.Engine{

	router := gin.Default()
	// 自定义渲染分隔符
	router.Delims("{[{", "}]}")
	// 使用日志中间件
	router.Use(gin.Logger())
	// 使用Recovery中间件
	router.Use(gin.Recovery())
	// 注册web路由
	WebRegister(router)
	// 注册api路由
	ApiRegister(router)
	return router
}
