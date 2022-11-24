package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"gorm.io/gorm"
)

func GetTask(db *gorm.DB, id int) (entities.Task, error) {
	var task entities.Task
	err := db.First(&task, id).Error
	return task, err
}

func GetTasks(db *gorm.DB, ids []int) ([]entities.Task, error) {
	var tasks []entities.Task
	err := db.Find(&tasks, ids).Error
	return tasks, err
}

func GetAllTasks(db *gorm.DB) ([]entities.Task, error) {
	var tasks []entities.Task
	err := db.Find(&tasks).Error
	return tasks, err
}

func CreateTask(db *gorm.DB, task *entities.Task) error {
	err := db.Create(task).Error
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
