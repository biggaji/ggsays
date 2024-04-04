package dtos

import (
	"github.com/biggaji/ggsays/models"
	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID    `json:"id"`
	Content   string       `json:"content"`
	UserId    uuid.UUID    `json:"userId"`
	LikeCount int32        `json:"favoriteCount"`
	User      *models.User `json:"author,omitempty"`
}
