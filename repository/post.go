package repository

import (
	"errors"

	"github.com/biggaji/ggsays/database"
	"github.com/biggaji/ggsays/helper"
	"github.com/biggaji/ggsays/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InsertNewPostRecord(post model.Post) error {
	result := database.Client.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPostById(postId uuid.UUID, includeUser bool) (model.Post, error) {
	var post model.Post

	db := database.Client.Model(&post).Where("id = ?", postId)
	if includeUser {
		db = db.Preload("User").Omit("password")
	}

	if err := db.Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return post, helper.ErrPostNotFound
		}
		return post, err
	}

	return post, nil
}

func GetAllPosts(includeUser bool) []model.Post {
	var posts []model.Post

	query := database.Client.Model(&posts)

	if includeUser {
		query = query.Preload("User").Omit("password")
	}

	query.Find(&posts)

	return posts
}

func PostContentExistForUser(userId uuid.UUID, content string) bool {
	var post model.Post
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
	// var post model.Post

	return nil
}
