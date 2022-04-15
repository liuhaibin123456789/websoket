package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"websoket/lv1/api"
	"websoket/lv1/middlewares"
	"websoket/lv1/tool"
)

func main() {
	err := tool.InitMysql()
	if err != nil {
		log.Println(err)
		panic(err)
		return
	}

	err = tool.InitRedis()
	if err != nil {
		log.Println(err)
		panic(err)
		return
	}
	go tool.FirstClientManager.Start()

	router := gin.Default()
	router.POST("register", api.Register)
	router.POST("login", api.Login)
	router.GET("chat", middlewares.CheckCookie, api.Chat)
	router.GET("message", middlewares.CheckCookie, api.GetMessage)
	router.Run()
}
