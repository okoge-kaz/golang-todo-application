package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/db"
	"github.com/okoge-kaz/golang-todo-application/server/entities"
)

func GetOwnership(userID int, taskID int) (entities.Ownership, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return entities.Ownership{}, err
	}

	var ownership entities.Ownership
	err = db.Where("user_id = ? AND task_id = ?", userID, taskID).First(&ownership).Error
	return ownership, err
}

func CreateOwnership(user entities.User, task entities.Task) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	ownership := entities.Ownership{
		UserID: user,
		TaskID: task,
	}

	err = db.Create(&ownership).Error
	return err
}

func UpdateOwnership(ownership *entities.Ownership) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Save(ownership).Error
	return err
}

func DeleteOwnership(ownership *entities.Ownership) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Delete(ownership).Error
	return err
}
