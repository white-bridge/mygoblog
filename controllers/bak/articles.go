package bak

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"asablog/models"
	"strconv"
)

type Articles struct {

}

func (a *Articles) Index(c *gin.Context)  {
	articleModel := new(models.Articles)
	articles := articleModel.List()
	c.HTML(http.StatusOK,"bak/articles.html",gin.H{"articles":articles})
}



func (a *Articles) Edit(c *gin.Context)  {
	articleModel := new(models.Articles)
	id := c.Query("id")
	if id != "" {
		id,_ := strconv.Atoi(id)
		articleModel.First(id)
	}
	c.HTML(http.StatusOK,"bak/articleEdit.html",gin.H{"article":articleModel})
}