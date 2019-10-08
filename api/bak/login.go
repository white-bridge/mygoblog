package bak

import (
	"asablog/models"
	"github.com/gin-gonic/gin"
	"crypto/md5"
	"strings"
	"encoding/hex"
	"asablog/info"
	"asablog/auth"
	"net/http"
	"encoding/json"
	"fmt"
)

// 后台用户登录
func Verify(c *gin.Context)  {

	var r models.Admin
	err := json.NewDecoder(c.Request.Body).Decode(&r)
	fmt.Println(err)
	fmt.Println(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"data":  "",
			"message": "bad params",
		})
		return
	}

	adminModel := new(models.Admin)
	admin := adminModel.First("name = ?",r.Name)

	if admin.ID != 0 {
		h := md5.New()
		h.Write([]byte(r.Password + admin.Salt))
		if strings.EqualFold(hex.EncodeToString(h.Sum(nil)), admin.Password) {
			// 登录成功
			token, _ := auth.GenerateBakToken(admin)
			info.AUser = append(info.AUser,info.AdminUser{Id:admin.ID,Name:admin.Name})
			c.JSON(http.StatusOK, gin.H{
				"code":  http.StatusOK,
				"data":  models.JwtToken{Token: token},
				"message": "登录成功",
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  http.StatusBadRequest,
				"data":  "",
				"message": "密码错误",
			})
		}
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"data":  "",
			"message": "用户不存在",
		})
	}
}
