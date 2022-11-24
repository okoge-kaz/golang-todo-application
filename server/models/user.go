package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/db"
	"github.com/okoge-kaz/golang-todo-application/server/entities"
)

func GetUserByID(userID int) (entities.User, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return entities.User{}, err
	}

	var user entities.User
	err = db.First(&user, userID).Error
	return user, err
}

func GetUsers(ids []int) ([]entities.User, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	var users []entities.User
	err = db.Find(&users, ids).Error
	return users, err
}

func GetAllUsers() ([]entities.User, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	var users []entities.User
	err = db.Find(&users).Error
	return users, err
}

func CreateUser(user *entities.User) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Create(user).Error
	return err
}

func UpdateUser(user *entities.User) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Save(user).Error
	return err
}

func DeleteUser(user *entities.User) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Delete(user).Error
	return err
}
