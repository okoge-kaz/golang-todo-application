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
	// SELECT * FROM ownerships WHERE user_id = userID AND task_id = taskID;
	return ownership, err
}

func CreateOwnership(user entities.User, task entities.Task) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	ownership := entities.Ownership{
		UserID: int(user.ID),
		TaskID: int(task.ID),
	}

	err = db.Create(&ownership).Error
	// INSERT INTO ownerships (user_id, task_id) VALUES (user.ID, task.ID);
	return err
}

func UpdateOwnership(ownership *entities.Ownership) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Save(ownership).Error
	// UPDATE ownerships SET user_id = ownership.UserID, task_id = ownership.TaskID WHERE id = ownership.ID;
	return err
}

func DeleteOwnership(ownership *entities.Ownership) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Delete(ownership).Error
	// DELETE FROM ownerships WHERE id = ownership.ID;
	return err
}
