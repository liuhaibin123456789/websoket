package model

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// Upgrade 协议升级
var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

// RoomMaker 房间创建者表
type RoomMaker struct {
	RoomId int `json:"room_id" gorm:"type:int;not null"`
	UserId int `json:"user_id" gorm:"type:int;not null"`
}

func (RoomMaker) TableName() string {
	return "room_maker"
}

// RoomManager 房间管理员表
type RoomManager struct {
	RoomId int `json:"room_id" gorm:"type:int;not null"`
	UserId int `json:"user_id" gorm:"type:int;not null"`
}

func (RoomManager) TableName() string {
	return "room_manager"
}

// RoomMember 房间普通成员
type RoomMember struct {
	RoomId int `json:"room_id" gorm:"type:int;not null"`
	UserId int `json:"user_id" gorm:"type:int;not null"`
}

func (RoomMember) TableName() string {
	return "room_member"
}

// RoomForbid 房间禁言成员
type RoomForbid struct {
	RoomId int `json:"room_id" gorm:"type:int;not null"`
	UserId int `json:"user_id" gorm:"type:int;not null"`
}

func (RoomForbid) TableName() string {
	return "room_forbid"
}
