package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `db:"title"`
	Status      string `db:"status"`
	Description string `db:"description"`
	IsDone      bool   `db:"is_done"`
	CategoryID  uint   `db:"category_id"`
}
