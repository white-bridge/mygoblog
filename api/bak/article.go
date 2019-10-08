package bak

import (
	"github.com/gin-gonic/gin"
	"asablog/models"
	"net/http"
	"fmt"
)

type Article struct {

}

func(a *Article) List()  {

}

func(a *Article) Edit( c *gin.Context)  {
	articleModel := new(models.Articles)
	err := c.BindJSON(articleModel)
	fmt.Println("err :",err)
	fmt.Println("id :",articleModel.Id)
	if articleModel.Id == 0 {
		articleModel.Insert(articleModel)
	} else {
		articleModel.Edit(articleModel)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"data":  "",
		"message": "保存成功",
	})
}