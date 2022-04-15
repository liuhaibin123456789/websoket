package dao

import (
	"errors"
	"log"
	"websoket/lv1/model"
	"websoket/lv1/tool"
)

func Register(user model.User) error {
	//事务
	tx := tool.GDB.Begin()

	defer func() {
		if r := recover; r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&user).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	//随即用户名
	if err := tx.Create(&model.UserSide{Name: tool.RandString(16)}).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func Login(user model.User) error {
	u := model.User{}
	if err := tool.GDB.Model(&model.User{}).Where("phone=?", user.Phone).Find(&u).Error; err != nil {
		log.Println(err)
		return err
	}

	if user.Password != u.Password {
		err := errors.New("password is wrong")
		log.Println(err)
		return err
	}
	return nil
}
