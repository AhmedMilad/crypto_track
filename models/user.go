package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FullName  string `gorm:"column:full_name;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
