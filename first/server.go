package main

import (
	"encoding/json"
	"log"
)

//ClientsManger 客户端管理者：负责收集所有客户端消息并选择分发；保存在线和非在线用户
type ClientsManger struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

var _ CMInterface = &ClientsManger{}

type CMInterface interface {
	Start() error
	Send(message []byte, ignoreClient *Client)
}

func (cm *ClientsManger) Send(message []byte, ignoreClient *Client) {
	for c := range cm.Clients {
		//分发没有忽略用户消息
		if c != ignoreClient {
			c.Message <- message
		}
	}
	return
}

func (cm *ClientsManger) Start() error {
	for true {

		select {

		case client := <-cm.Register:
			//注册用户塞进用户管理器
			cm.Clients[client] = true
			msg, err := json.Marshal(&Message{Content: "a new socket is open..."})
			if err != nil {
				log.Println(err)
				return err
			}
			//发送给其他客户端
			cm.Send(msg, client)

		case client := <-cm.Unregister:
			//关闭管道
			close(client.Message)
			//客户端管理器删掉注销用户
			delete(cm.Clients, client)
			msg, err := json.Marshal(&Message{Content: "a socket is closed..."})
			if err != nil {
				log.Println(err)
				return err
			}
			//发送给其他客户端
			cm.Send(msg, client)

		case broadcast := <-cm.Broadcast:
			//遍历连接的客户端发送消息
			for client := range cm.Clients {
				select {
				case client.Message <- broadcast:
				default:
					close(client.Message)
					delete(cm.Clients, client)
				}
			}
			cm.Send(broadcast, nil)
		}
	}
	return nil
}

// NewClientsManager 创建客户端管理器
func newClientsManager() *ClientsManger {
	clientsManger := &ClientsManger{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
	//较大的结构体传指针更有效率
	return clientsManger
}
