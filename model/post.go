package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primarykey"`
	Content   string         `json:"content" gorm:"not null;index"`
	UserId    uuid.UUID      `json:"userId" gorm:"not null;index"`
	User      User           `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
	LikeCount int32          `json:"favoriteCount" gorm:"default:0"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime,default:current_timestamp"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
