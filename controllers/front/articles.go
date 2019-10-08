package front

import (
	"asablog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/microcosm-cc/bluemonday"
)

type Articles struct {
	
}

func (a *Articles) Index(c *gin.Context)  {
	articleModel := new(models.Articles)
	page,_:=strconv.Atoi(c.Query("page"))
	stype,_:=strconv.Atoi(c.Query("type"))
	if page == 0 {
		page = 1
	}
	articles := articleModel.ListCondition(10,10*(page-1),map[string]interface{}{"check_status": 2, "type": stype,"isdel":1})
	// 统计总页数
	rows := articleModel.CountRows(map[string]interface{}{"check_status": 2, "type": stype,"isdel":1})
	totalPage := rows / 10
	if rows % 10 != 0 {
		totalPage += 1
	}
	prev := page - 1
	next := page + 1
	if totalPage == 1 {
		prev = 1
		next = 1
	} else {
		if prev <= 0 {
			prev = 1
		}
		if next >= totalPage {
			next = totalPage
		}
	}
	c.HTML(http.StatusOK,"front/articles.html",gin.H{
		"articles":articles,
		"page":page,
		"type":stype,
		"prev":prev,
		"next":next,
		"totalPage":totalPage,
	})
}

func (a *Articles) Single(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Query("id"))
	articleModel := new(models.Articles)
	article := articleModel.First(id)
	unsafe := blackfriday.Run([]byte(article.Content))
	article.Content = string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
	c.HTML(http.StatusOK,"front/single.html",gin.H{"article":article})
}