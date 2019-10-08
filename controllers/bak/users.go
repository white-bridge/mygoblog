package bak

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {

}

func (u *User) Index(c *gin.Context)  {

	c.HTML(http.StatusOK,"bak/users.html",gin.H{})
}