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
		//注意http请求被劫持，无法保障原有的http请求的有效性，所以使用http对应的方法可能会报错
		return
	}
	//协议升级
	conn, err := model.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	err = service.Chat(phone.(string), conn)
	if err != nil {
		log.Println(err)
		return
	}
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
