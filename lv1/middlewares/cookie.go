package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"websoket/lv1/tool"
)

func CheckCookie(c *gin.Context) {
	defer func() {
		phone, err := c.Cookie("phone")
		if err != nil {
			log.Println(err)
			tool.JsonOut(c, err, "cookie失效，请重新登录")
			c.Abort()
			return
		}
		if len(phone) != 11 {
			tool.JsonOut(c, errors.New("cookie error"), "cookie错误")
			c.Abort()
			return
		}
		c.Set("phone", phone)
		c.Next()
	}()
}
