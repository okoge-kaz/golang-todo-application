package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, id int) (entities.User, error) {
	var user entities.User
	err := db.First(&user, id).Error
	return user, err
}

func GetUsers(db *gorm.DB, ids []int) ([]entities.User, error) {
	var users []entities.User
	err := db.Find(&users, ids).Error
	return users, err
}

func GetAllUsers(db *gorm.DB) ([]entities.User, error) {
	var users []entities.User
	err := db.Find(&users).Error
	return users, err
}

func CreateUser(db *gorm.DB, user *entities.User) error {
	err := db.Create(user).Error
	return err
}

func UpdateUser(db *gorm.DB, user *entities.User) error {
	err := db.Save(user).Error
	return err
}

func DeleteUser(db *gorm.DB, id int) error {
	err := db.Delete(&entities.User{}, id).Error
	return err
}

func DeleteUsers(db *gorm.DB, ids []int) error {
	err := db.Delete(&entities.User{}, ids).Error
	return err
}
