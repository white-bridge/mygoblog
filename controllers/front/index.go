package front

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"asablog/models"
	"strconv"
)

func Index(c *gin.Context)  {
	articleModel := new(models.Articles)
	page,_:=strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	articles := articleModel.ListCondition(10,10*(page-1),map[string]interface{}{"check_status": 2,"isdel":1})
	// 统计总页数
	rows := articleModel.CountRows(map[string]interface{}{"check_status": 2,"isdel":1})
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

	c.HTML(http.StatusOK,"front/index.html",gin.H{
		"articles":articles,
		"page":page,
		"prev":prev,
		"next":next,
		"totalPage":totalPage,
	})
}

//关于我
func About(c *gin.Context){

	c.HTML(http.StatusOK,"front/about.html",gin.H{})
}