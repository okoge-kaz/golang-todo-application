package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/db"
	"github.com/okoge-kaz/golang-todo-application/server/entities"
)

func GetCategoryByID(categoryID int) (entities.Category, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return entities.Category{}, err
	}

	var category entities.Category
	err = db.First(&category, categoryID).Error
	// SELECT * FROM categories WHERE id = categoryID
	return category, err
}

func GetCategories(ids []int) ([]entities.Category, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	var categories []entities.Category
	err = db.Find(&categories, ids).Error
	// SELECT * FROM categories WHERE id IN ids
	return categories, err
}

func GetAllCategories() ([]entities.Category, error) {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return nil, err
	}

	var categories []entities.Category
	err = db.Find(&categories).Error
	// SELECT * FROM categories
	return categories, err
}

func DeleteCategory(category *entities.Category) error {
	// connect to database
	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	err = db.Delete(category).Error
	// DELETE FROM categories WHERE id = category.ID
	return err
}
