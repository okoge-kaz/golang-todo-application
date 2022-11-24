package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"gorm.io/gorm"
)

func GetCategory(db *gorm.DB, id int) (entities.Category, error) {
	var category entities.Category
	err := db.First(&category, id).Error
	return category, err
}

func GetCategories(db *gorm.DB, ids []int) ([]entities.Category, error) {
	var categories []entities.Category
	err := db.Find(&categories, ids).Error
	return categories, err
}

func GetAllCategories(db *gorm.DB) ([]entities.Category, error) {
	var categories []entities.Category
	err := db.Find(&categories).Error
	return categories, err
}

func CreateCategory(db *gorm.DB, category *entities.Category) error {
	err := db.Create(category).Error
	return err
}

func UpdateCategory(db *gorm.DB, category *entities.Category) error {
	err := db.Save(category).Error
	return err
}

func DeleteCategory(db *gorm.DB, id int) error {
	err := db.Delete(&entities.Category{}, id).Error
	return err
}

func DeleteCategories(db *gorm.DB, ids []int) error {
	err := db.Delete(&entities.Category{}, ids).Error
	return err
}
