package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/db"
	"github.com/okoge-kaz/golang-todo-application/server/entities"
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
	// INSERT INTO tasks (title, description, deadline, status) VALUES (task.Title, task.Description, task.Deadline, task.Status)
	return err
}

// Transaction
func CreateTaskWithTransaction(task *entities.Task, user entities.User) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	// start transaction
	tx := db.Begin()
	err = tx.Create(task).Error
	// INSERT INTO tasks (title, description, deadline, status) VALUES (task.Title, task.Description, task.Deadline, task.Status)
	if err != nil {
		tx.Rollback()
		return err
	}

	ownership := entities.Ownership{
		UserID: int(user.ID),
		TaskID: int(task.ID),
	}
	err = tx.Create(&ownership).Error
	// INSERT INTO ownerships (user_id, task_id) VALUES (user.ID, task.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// commit transaction
	tx.Commit()
	return nil
}

func UpdateTask(task entities.Task, updateTask entities.Task) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Model(&task).Updates(updateTask).Error // UPDATE tasks SET title = updateTask.Title, description = updateTask.Description, deadline = updateTask.Deadline, status = updateTask.Status WHERE id = task.ID
	return err
}

func DeleteTask(taskID int) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Delete(&entities.Task{}, taskID).Error // DELETE FROM tasks WHERE id = taskID
	return err
}

func DeleteTasks(ids []int) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Delete(&entities.Task{}, ids).Error // DELETE FROM tasks WHERE id IN (ids)
	return err
}
