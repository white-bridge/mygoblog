package front

import (
	"asablog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Articles struct {
	
}

func (a *Articles) Index(c *gin.Context)  {
	articleModel := new(models.Articles)
	articles := articleModel.List()
	c.HTML(http.StatusOK,"front/articles.html",gin.H{"articles":articles})
}