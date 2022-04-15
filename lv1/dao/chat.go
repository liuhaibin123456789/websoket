package dao

import (
	"log"
	model "websoket/lv1/model"
	"websoket/lv1/tool"
)

func SelectUser(client *tool.Client, phone string) error {
	user := &model.UserSide{}
	var id int

	if err := tool.GDB.Model(&model.User{}).Select("id").Where("phone=?", phone).Find(&id).Error; err != nil {
		log.Println(err)
		return err
	}

	if err := tool.GDB.Model(user).Where("id=?", id).Find(&user).Error; err != nil {
		log.Println(err)
		return err
	}

	//填充client信息
	client.Id = user.Id
	client.Sex = user.Sex
	client.UserName = user.Name

	return nil
}

func SelectMessage(phone string) ([]model.Message, error) {
	//查找发送消息的用户id
	var id int
	if err := tool.GDB.Model(&model.User{}).Select("id").Where("phone=?", phone).Find(&id).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	messages := make([]model.Message, 0)

	if err := tool.GDB.Model(&model.Message{}).Where("chat_room_id=? and sender_id=?", 1, id).Find(&messages).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return messages, nil
}
