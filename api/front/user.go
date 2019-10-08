package front

import (
	"net/http"
	"asablog/models"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"asablog/auth"
	"crypto/md5"
	"strings"
	"encoding/hex"
)

// 前台用户注册
func Register(c *gin.Context) {
	r := c.Request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Nickname == "" || user.Password == "" || user.Email == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"data":  "",
			"message": "bad params",
		})
		return
	}
	//判断邮箱是否已存在
	old := user.First("email = ?",user.Email)
	if old.ID != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"data":  "",
			"message": "email already exsit",
		})
		return
	}
	id := user.Insert(&user)
	if id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  http.StatusInternalServerError,
			"data":  "",
			"message": "internal error",
		})
	}
	token, _ := auth.GenerateToken(&user)
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"data":  models.JwtToken{Token: token},
		"message": "",
	})
}

// 前台用户登录
func Login(c *gin.Context) {
	r := c.Request
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"data":  "",
			"message": "bad params",
		})
		return
	}
	//先保存密码
	pw := user.Password
	user.First("email = ?",user.Email)
	if user.ID != 0 {
		h := md5.New()
		h.Write([]byte(pw + user.Salt))
		if strings.EqualFold(hex.EncodeToString(h.Sum(nil)), user.Password) {
			token, _ := auth.GenerateToken(&user)
			c.JSON(http.StatusOK, gin.H{
				"code":  http.StatusOK,
				"data":  models.JwtToken{Token: token},
				"message": "",
			})
		}else{
			c.JSON(http.StatusNotFound, gin.H{
				"code":  http.StatusNotFound,
				"data":  "",
				"message": "password wrong",
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"data":  "",
			"message": "the user not exist",
		})
	}
}

