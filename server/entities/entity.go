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
	IsDone      bool      `gprm:"notNull:IsDone"`        // tasks.is_done // TODO Not NULL Constraint
	CategoryID  Category  `gorm:"foreignKey:CategoryID"` // tasks.category_id (has one)
}

type User struct {
	gorm.Model
	Name     string // users.name
	Password []byte // users.password
}

type Ownership struct {
	gorm.Model
	UserID int
	User   User `gorm:"foreignKey:UserID"`
	TaskID int
	Task   Task `gorm:"foreignKey:TaskID"`
}

type Category struct {
	gorm.Model
	Name  string // categories.name
	Color string // categories.color
}
