package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/db"
	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"github.com/okoge-kaz/golang-todo-application/server/helpers"
)

func GetUserByID(userID int) (entities.User, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return entities.User{}, err
	}

	var user entities.User
	err = db.First(&user, userID).Error
	// SELECT * FROM users WHERE id = userID;
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
	// SELECT * FROM users WHERE id IN (ids);
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
	// SELECT * FROM users;
	return users, err
}

func CreateUser(user *entities.User) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Create(user).Error
	// INSERT INTO users (name, user_name, password) VALUES (user.Name, user.UserName, user.Password);
	return err
}

func CheckDuplicateUserName(userName string) bool {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return false
	}

	var count int64
	db.Model(&entities.User{}).Where("user_name = ?", userName).Count(&count)
	// SELECT count(*) FROM users WHERE user_name = 'userName';
	return count > 0
	// return true if user name is already taken
}

func UpdateUser(user *entities.User) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Save(user).Error
	// UPDATE users SET name = user.Name, password = user.Password WHERE id = user.ID;
	return err
}

func ChangeUserPassword(user *entities.User, password string) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	user.Password = helpers.EncryptPassword(password)
	err = db.Save(user).Error
	// UPDATE users SET password = password WHERE id = user.ID;
	return err
}

func ChangeUserName(user *entities.User, userName string) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	user.Name = userName
	err = db.Save(user).Error
	// UPDATE users SET name = userName WHERE id = user.ID;
	return err
}

func DeleteUser(user *entities.User) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Delete(user).Error
	// DELETE FROM users WHERE id = user.ID;
	return err
}
