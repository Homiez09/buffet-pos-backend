package models

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Manager  Role = "manager"
	Employee Role = "employee"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	Role      Role      `json:"role" gorm:"type:varchar(50);default:'employee'"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
