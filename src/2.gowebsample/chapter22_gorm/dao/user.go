package dao

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	UserName string `gorm:"column:user_name"`
	Passwd   string `gorm:"column:passwd" `
	Avatar   string `gorm:"column:avatar" `
}

func (User) TableName() string {
	return "gorm_user"
}

func SaveUser(user *User) (User, error) {
	err := DB.Create(user).Error
	if err != nil {
		log.Fatalf("save user err:%v", err)
		return User{}, err
	}
	return *user, err
}

func GetUser(id int) (*User, error) {
	var user User
	err := DB.Model(user).First(&user, id).Error
	if err != nil {
		log.Fatalf("save user err:%v", err)
		return nil, err
	}
	return &user, err
}

func GetUserAll() ([]*User, error) {
	var user []*User
	err := DB.Find(&user).Error
	if err != nil {
		log.Fatalf("save user err:%v", err)
		return nil, err
	}
	return user, err
}

func UpdateUser(user *User) error {

	err := DB.Model(&user).Where("id=?", user.ID).Updates(User{
		UserName: user.UserName,
		Passwd:   user.Passwd,
		Avatar:   user.Avatar,
	}).Error

	if err != nil {
		log.Fatalf("save user err:%v", err)
		return err
	}
	return nil
}

func DeleteUser(user *User) error {

	err := DB.Where("id=?", user.ID).Delete(&user).Error
	if err != nil {
		log.Fatalf("save user err:%v", err)
		return err
	}
	return nil
}
