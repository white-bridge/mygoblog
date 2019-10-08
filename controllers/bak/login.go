package bak

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {

	c.HTML(http.StatusOK,"bak/login.html",gin.H{})
}
