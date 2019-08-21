package routers

import "github.com/gin-gonic/gin"

func Register() *gin.Engine{
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	WebRegister(router)
	ApiRegister(router)
	return router
}
