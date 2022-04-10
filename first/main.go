package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"log"
)

// CM 初始化客户端管理器
var CM = newClientsManager()

func main() {
	//开一个议程管理客户端注册注销与广播
	go CM.Start()

	//注册web服务监听
	router := gin.Default()

	router.GET("/ws", ws)
	router.Run()
}

//使用websocket协议，需要将http协议经过第三方框架处理升级，添加websocket协议专属字段
func ws(c *gin.Context) {
	//用户名
	userNmae := c.DefaultQuery("user_name", "un")

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	//协议升级，拿到websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//开启客户端
	client := &Client{
		Id:       uuid.Must(uuid.NewV4(), nil).String(),
		UserName: userNmae,
		Message:  make(chan []byte),
		Conn:     conn,
	}

	//注册客户端
	CM.Register <- client

	//启动协程收web端传过来的消息
	go client.Write(CM)
	//启动协程把消息返回给web端
	go client.Read()

}
