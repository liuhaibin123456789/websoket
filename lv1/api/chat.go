package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"websoket/lv1/model"
	"websoket/lv1/service"
	"websoket/lv1/tool"
)

func Chat(c *gin.Context) {
	phone, res := c.Get("phone")
	if !res {
		err := errors.New("phone is not got")
		log.Println(err)
		tool.JsonOut(c, err, "cookie 获取失败")
		return
	}
	//协议升级
	conn, err := model.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		tool.JsonOut(c, err, "协议升级失败")
		return
	}
	err = service.Chat(phone.(string), conn)
	if err != nil {
		log.Println(err)
		tool.JsonOut(c, err, "聊天创建失败")
		return
	}
	tool.JsonOut(c, err, "欢迎进入聊天室")
}

func GetMessage(c *gin.Context) {
	phone := c.GetString("phone")
	messages, err := service.GetMessage(phone)
	if err != nil {
		tool.JsonOut(c, err, "获取历史消息失败")
		return
	}
	tool.JsonOut(c, nil, messages)
}
