package repository

import (
	"errors"

	"github.com/biggaji/ggsays/database"
	"github.com/biggaji/ggsays/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InsertNewPostRecord(post models.Post) error {
	result := database.Client.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPostById(postId uuid.UUID, includeUser bool) (models.Post, error) {
	var post models.Post

	db := database.Client.Model(&post).Where("id = ?", postId)
	if includeUser {
		db = db.Preload("User").Omit("password")
	}

	if err := db.Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return post, errors.New("post not found")
		}
		return post, err
	}

	return post, nil
}

func GetAllPosts(includeUser bool) []models.Post {
	var posts []models.Post

	query := database.Client.Model(&posts)

	if includeUser {
		query = query.Preload("User").Omit("password")
	}

	query.Find(&posts)

	return posts
}

func PostContentExistForUser(userId uuid.UUID, content string) bool {
	var post models.Post
	result := database.Client.Where("user_id = ?", userId).Where("LOWER(content) = LOWER(?)", content).First(&post)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func DeletePost() error {
	return nil
}

func PurgePosts() error {
	return nil
}

func UpvotePost() error {
	return nil
}
