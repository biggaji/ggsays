package dto

import (
	"github.com/biggaji/ggsays/model"
	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID   `json:"id"`
	Content   string      `json:"content"`
	UserId    uuid.UUID   `json:"userId"`
	LikeCount int32       `json:"favoriteCount"`
	User      *model.User `json:"author,omitempty"`
}
