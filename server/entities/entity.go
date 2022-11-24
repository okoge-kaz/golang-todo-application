package entities

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    // tasks.title
	Deadline    time.Time // tasks.deadline
	Status      string    // tasks.status
	Description string    // tasks.description
	IsDone      bool      // tasks.is_done
	CategoryID  Category  `gorm:"foreignKey:CategoryID"` // tasks.category_id
}

type User struct {
	gorm.Model
	Name     string // users.name
	Password []byte // users.password
}

type Ownership struct {
	gorm.Model
	UserID User `gorm:"foreignKey:UserID"` // ownerships.user_id // belongs to user (many to one) (many ownerships can belong to one user)
	TaskID Task `gorm:"foreignKey:TaskID"` // ownerships.task_id // has one task (one to one) (one ownership has one task)
}

type Category struct {
	gorm.Model
	Name string // categories.name
}
