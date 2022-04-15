package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"websoket/lv1/model"
	"websoket/lv1/service"
	"websoket/lv1/tool"
)

func Register(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	err := service.Register(model.User{Phone: phone, Password: password})
	if err != nil {
		log.Println(err)
		tool.JsonOut(c, err, "注册失败")
		return
	}
	//生成cookie
	c.SetCookie("phone", phone, 1000, "/", "", true, false)
	tool.JsonOut(c, nil, "注册成功")
}

func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	err := service.Login(model.User{Phone: phone, Password: password})
	if err != nil {
		log.Println(err)
		tool.JsonOut(c, err, "登录失败")
		return
	}
	//生成cookie
	c.SetCookie("phone", phone, 1000, "/", "", true, false)
	tool.JsonOut(c, nil, "登录成功")
}
