package model

type Message struct {
	//标识消息的唯一性
	Id int `json:"id" gorm:"primaryKey;autoIncrement"`
	//聊天室id
	ChatRoomId int `json:"chat_room_id" gorm:"type:int;not null"`
	//发表用户的id
	SenderId int `json:"sender_id" gorm:"type:int;not null"`
	//内容
	Content string `json:"content" gorm:"type:varchar(1000);not null"` //限制发表字符数最多为1000字符
}

func (m Message) TableName() string {
	return "message"
}
