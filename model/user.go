package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primarykey"`
	FirstName string         `json:"firstName" gorm:"not null;index"`
	LastName  string         `json:"lastName" gorm:"not null;index"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Password  string         `json:"password" gorm:"not null"`
	UserName  string         `json:"userName" gorm:"not null;unique;index"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
