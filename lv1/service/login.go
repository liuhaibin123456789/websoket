package service

import (
	"errors"
	"log"
	"websoket/lv1/dao"
	"websoket/lv1/model"
)

func Register(user model.User) error {
	//简单校验
	if len(user.Phone) != 11 {
		err := errors.New("the phone is wrong")
		if err != nil {
			log.Println(err)
			return err
		}
	}
	//限制密码长度
	if length := len(user.Password); length < 8 || length > 16 {
		err := errors.New("the password is wrong")
		if err != nil {
			log.Println(err)
			return err
		}
	}

	err := dao.Register(user)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func Login(user model.User) error {
	//简单校验
	if len(user.Phone) != 11 {
		err := errors.New("the phone is wrong")
		if err != nil {
			log.Println(err)
			return err
		}
	}

	if length := len(user.Password); length < 8 || length > 16 {
		err := errors.New("the password is wrong")
		if err != nil {
			log.Println(err)
			return err
		}
	}
	err := dao.Login(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
