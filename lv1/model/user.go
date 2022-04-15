package model

// UserSide 用户非隐私数据表
type UserSide struct {
	Id   int    `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(20);not null"`
	Sex  string `json:"sex" gorm:"type:varchar(5);not null"`
}

func (u UserSide) TableName() string {
	return "user_side"
}

// User 用户隐私数据表
type User struct {
	Id       int    `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Phone    string `json:"phone" gorm:"type:varchar(11);not null"`
	Password string `json:"password" gorm:"type:varchar(16);not null"`
}

func (u User) TableName() string {
	return "user"
}
