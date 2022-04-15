package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"websoket/lv1/model"
)

type Client struct {
	Id       int             `json:"id"`
	UserName string          `json:"userName"`
	Sex      string          `json:"sex"`
	Message  chan []byte     `json:"message"`
	Conn     *websocket.Conn `json:"-"`
}

//Write 将消息读取出来,写入链接里
func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()
	ticker := time.NewTicker(55 * time.Second)
	for true {
		select {
		case message, ok := <-c.Message:
			if !ok {
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					fmt.Println(err)
					return
				}
			}
			//写入客户端
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Println(err)
				return
			}
		//心跳,处理ping
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
	return
}

//Read 广播消息
func (c *Client) Read(cm *ClientManager) {
	defer func() {
		cm.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil })
	for true {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		bytes1, err := json.Marshal(model.Message{SenderId: c.Id, Content: string(message)})
		if err != nil {
			log.Println(err)
			return
		}
		//去掉json的空闲格式，压缩数据大小
		bytes1 = bytes.TrimSpace(bytes.Replace(bytes1, []byte("\n"), []byte(" "), -1))
		cm.Send(bytes1, c)
	}
}

// ClientManager 客户端管理器
type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

var FirstClientManager = NewClientManager(50)

func NewClientManager(cap int) *ClientManager {
	//创建新的客户端管理器
	return &ClientManager{
		Clients:    make(map[*Client]bool, cap),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (cm *ClientManager) Start() {
	for true {
		select {
		case client := <-cm.Register:
			//注册用户
			cm.Clients[client] = true
			//数据库操作：将用户存进对应普通成员表
			var rm = &model.RoomMember{
				RoomId: 1,
				UserId: client.Id,
			}
			if err := GDB.Model(&model.RoomMember{}).Create(rm).Error; err != nil {
				log.Println(err)
				return
			}
			bytes1, err := json.Marshal(&model.Message{SenderId: client.Id, ChatRoomId: 1, Content: client.UserName + "进入聊天室"})
			if err != nil {
				log.Println(err)
				return
			}
			bytes1 = bytes.TrimSpace(bytes.Replace(bytes1, []byte("\n"), []byte(" "), -1))
			cm.Send(bytes1, client)
			fmt.Println("进入聊天室")
		case client := <-cm.Unregister:
			//关闭管道
			close(client.Message)
			//内存删除
			delete(cm.Clients, client)
			//数据库操作
			var rm = &model.RoomMember{
				RoomId: 1,
				UserId: client.Id,
			}
			if err := GDB.Model(&model.RoomMember{}).Where("user_id=?", client.Id).Delete(rm).Error; err != nil {
				log.Println(err)
				return
			}
			bytes1, err := json.Marshal(&model.Message{SenderId: client.Id, ChatRoomId: 1, Content: client.UserName + "离开聊天室"})
			if err != nil {
				log.Println(err)
				return
			}
			bytes1 = bytes.TrimSpace(bytes.Replace(bytes1, []byte("\n"), []byte(" "), -1))
			cm.Send(bytes1, client)
			fmt.Println("离开聊天室")
		case broadcast := <-cm.Broadcast:
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
}

func (cm *ClientManager) Send(message []byte, ignoreClient *Client) {
	//数据库保存消息
	if err := GDB.Model(&model.Message{}).Create(&model.Message{SenderId: ignoreClient.Id, ChatRoomId: 1, Content: string(message)}).Error; err != nil {
		log.Println(err)
		return
	}
	for client := range cm.Clients {
		if client != ignoreClient {
			client.Message <- message
		}
	}
}
