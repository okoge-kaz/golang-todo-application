package entities

import "gorm.io/gorm"

type Ownership struct {
	gorm.Model
	UserID    uint `db:"user_id"`
	TaskID    uint `db:"task_id"`
}
