package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/db"
	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"gorm.io/gorm"
)

func GetTask(taskID int) (entities.Task, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return entities.Task{}, err
	}

	var task entities.Task
	err = db.First(&task, taskID).Error
	return task, err
}

func GetTasksByUserID(userID int, order string) ([]entities.Task, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	var tasks []entities.Task

	switch order {
	case "desc":
		err := db.Joins("JOIN ownerships ON ownerships.task_id = tasks.id").Where("ownerships.user_id = ?", userID).Order("deadline desc").Find(&tasks).Error
		return tasks, err
	default:
		err := db.Joins("JOIN ownerships ON ownerships.task_id = tasks.id").Where("ownerships.user_id = ?", userID).Order("deadline asc").Find(&tasks).Error
		return tasks, err
	}
}

func GetTasksByUserIDAndStatusAndKeyword(userID int, status []string, keywords []string, order string) ([]entities.Task, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	var tasks []entities.Task

	keywordsQuery := ""
	for i, keyword := range keywords {
		if i == 0 {
			keywordsQuery += "(title LIKE '%" + keyword + "%'"
		} else {
			keywordsQuery += " OR title LIKE '%" + keyword + "%'"
		}
	}
	keywordsQuery += ")"

	switch order {
	case "desc":
		err := db.Joins("JOIN ownerships ON ownerships.task_id = tasks.id").Where("ownerships.user_id = ? AND status IN (?) AND "+keywordsQuery, userID, status).Order("deadline desc").Find(&tasks).Error
		return tasks, err
	default:
		err := db.Joins("JOIN ownerships ON ownerships.task_id = tasks.id").Where("ownerships.user_id = ? AND status IN (?) AND "+keywordsQuery, userID, status).Order("deadline asc").Find(&tasks).Error
		return tasks, err
	}
}

func CreateTask(task *entities.Task) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Create(task).Error
	return err
}

func UpdateTask(db *gorm.DB, task *entities.Task) error {
	err := db.Save(task).Error
	return err
}

func DeleteTask(db *gorm.DB, id int) error {
	err := db.Delete(&entities.Task{}, id).Error
	return err
}

func DeleteTasks(db *gorm.DB, ids []int) error {
	err := db.Delete(&entities.Task{}, ids).Error
	return err
}
