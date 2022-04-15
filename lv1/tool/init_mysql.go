package tool

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"websoket/lv1/model"
)

// GDB 全局数据库操作对象
var GDB *gorm.DB

func InitMysql() error {
	//连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/websocket?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//默认关闭gorm的事务,提升性能
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}

	GDB = db

	//迁移表的数据
	if !GDB.Migrator().HasTable(&model.User{}) {
		err := GDB.AutoMigrate(&model.User{})
		if err != nil {
			return err
		}
	}

	if !GDB.Migrator().HasTable(&model.UserSide{}) {
		err := GDB.AutoMigrate(&model.UserSide{})
		if err != nil {
			return err
		}
	}

	if !GDB.Migrator().HasTable(&model.RoomMaker{}) {
		err := GDB.AutoMigrate(&model.RoomMaker{})
		if err != nil {
			return err
		}
	}

	if !GDB.Migrator().HasTable(&model.RoomManager{}) {
		err := GDB.AutoMigrate(&model.RoomManager{})
		if err != nil {
			return err
		}
	}

	if !GDB.Migrator().HasTable(&model.RoomMember{}) {
		err := GDB.AutoMigrate(&model.RoomMember{})
		if err != nil {
			return err
		}
	}
	//该表存入禁言时间>1小时
	if !GDB.Migrator().HasTable(&model.RoomForbid{}) {
		err := GDB.AutoMigrate(&model.RoomForbid{})
		if err != nil {
			return err
		}
	}
	if !GDB.Migrator().HasTable(&model.Message{}) {
		err := GDB.AutoMigrate(&model.Message{})
		if err != nil {
			return err
		}
	}

	return nil
}
