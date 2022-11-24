package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `db:"name"`
}
