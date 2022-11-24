package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `db:"name"`
	Password []byte `db:"password"`
}
