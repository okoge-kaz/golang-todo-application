package models

import (
	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"gorm.io/gorm"
)

func GetOwnership(db *gorm.DB, id int) (entities.Ownership, error) {
	var ownership entities.Ownership
	err := db.First(&ownership, id).Error
	return ownership, err
}

func GetOwnerships(db *gorm.DB, ids []int) ([]entities.Ownership, error) {
	var ownerships []entities.Ownership
	err := db.Find(&ownerships, ids).Error
	return ownerships, err
}

func GetAllOwnerships(db *gorm.DB) ([]entities.Ownership, error) {
	var ownerships []entities.Ownership
	err := db.Find(&ownerships).Error
	return ownerships, err
}

func CreateOwnership(db *gorm.DB, ownership *entities.Ownership) error {
	err := db.Create(ownership).Error
	return err
}

func UpdateOwnership(db *gorm.DB, ownership *entities.Ownership) error {
	err := db.Save(ownership).Error
	return err
}

func DeleteOwnership(db *gorm.DB, id int) error {
	err := db.Delete(&entities.Ownership{}, id).Error
	return err
}

func DeleteOwnerships(db *gorm.DB, ids []int) error {
	err := db.Delete(&entities.Ownership{}, ids).Error
	return err
}
