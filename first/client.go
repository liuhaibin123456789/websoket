package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Id       string          `json:"id"`
	UserName string          `json:"user_name"`
	Message  chan []byte     `json:"message"` //客户端发送的消息
	Conn     *websocket.Conn `json:"conn"`    //连接的socket
}

//静态校验方法是否实现完毕
var _ CInterface = &Client{}

type CInterface interface {
	// Write 往服务端写入广播消息
	Write(*ClientsManger) error

	// Read 从客户端的连接里读取消息
	Read()
}

func (c *Client) Write(cm *ClientsManger) error {
	//遇到panic或执行结束，需要注销此次websocket
	defer func() {
		cm.Unregister <- c
		c.Conn.Close()
	}()

	for true {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			//此处调用defer
			return err
			//cm.Unregister <- c
			//c.Conn.Close()
			//break
		}
		bytes, err := json.Marshal(&Message{Content: string(msg), Sender: c.Id})
		if err != nil {
			log.Println(err)
			//此处调用defer
			return err
		}
		//发送给其他客户端广播
		cm.Send(bytes, c)
	}
	return nil
}

func (c *Client) Read() {
	//读取结束或遇到panic，关闭连接
	defer func() {
		c.Conn.Close()
	}()

	for true {
		select {
		case msg, ok := <-c.Message:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//有消息就写入，发送给web端
			c.Conn.WriteMessage(websocket.TextMessage, msg)
		}
	}
	return
}
