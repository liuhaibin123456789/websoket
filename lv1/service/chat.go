package service

import (
	"github.com/gorilla/websocket"
	"log"
	"websoket/lv1/dao"
	"websoket/lv1/model"
	"websoket/lv1/tool"
)

func Chat(phone string, conn *websocket.Conn) error {
	client := &tool.Client{}
	//查询用户信息
	err := dao.SelectUser(client, phone)
	if err != nil {
		log.Println(err)
		return err
	}
	client.Message <- []byte(client.UserName + "加入聊天室")
	client.Conn = conn
	//注册客户端
	tool.FirstClientManager.Register <- client
	//启动一个协程,将该客户端的消息广播
	go client.Read(tool.FirstClientManager)
	//启动一个协程，从客户端与服务端的连接读取消息
	go client.Write()
	//多个协程同时读取数据库时，数据库引擎默认为innodb，为行锁。因此并发性能较好
	return nil
}

func GetMessage(phone string) ([]model.Message, error) {
	messages, err := dao.SelectMessage(phone)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return messages, nil
}
