package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

// 后台登录token验证
func ALogin()  gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("atoken")
		if tokenStr == "" {
			tokenStr,_ = c.Cookie("atoken")
		}
		if tokenStr == "" {
			c.Abort()
			c.Redirect(http.StatusFound, "/admin/login")
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					c.Abort()
					c.Redirect(http.StatusFound, "/admin/login")
				}
				return []byte("baksecret"), nil
			})
			if !token.Valid {
				c.Abort()
				c.Redirect(http.StatusFound, "/admin/login")
			} else {
				c.Next()
			}
		}
		c.Next()
	}
}

// 前台token验证

func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("authorization")
		if tokenStr == "" {
			tokenStr,_ = c.Cookie("authorization")
		}
		if tokenStr == "" {
			c.Abort()
			c.Redirect(http.StatusFound, "/")
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					c.Abort()
					c.Redirect(http.StatusFound, "/")
				}
				return []byte("frontsecret"), nil
			})
			if !token.Valid {
				c.Abort()
				c.Redirect(http.StatusFound, "/")
			} else {
				c.Next()
			}
		}
	}
}
