package bak

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK,"bak/index.html",gin.H{})
}