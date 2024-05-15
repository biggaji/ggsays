package repository

import (
	"errors"
	"fmt"

	"github.com/biggaji/ggsays/database"
	"github.com/biggaji/ggsays/model"
	"gorm.io/gorm"
)

func InsertUserRecord(user model.User) error {
	result := database.Client.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserById(id uint) (model.User, error) {
	var user model.User
	result := database.Client.Take(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user with ID %d not found", id)
		}
		return user, result.Error
	}
	return user, nil
}

func UserRecordExist(identifier string) bool {
	var user model.User
	result := database.Client.First(&user, "email = ? OR user_name = ?", identifier, identifier)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	result := database.Client.Take(&user, "email = ?", email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, fmt.Errorf("user with email %s not found", email)
		}
		return model.User{}, fmt.Errorf("failed to get user by email: %v", result.Error)
	}

	return user, nil
}

func GetUserByUsername(userName string) (model.User, error) {
	var user model.User
	result := database.Client.Take(&user, "user_name = ?", userName)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.User{}, fmt.Errorf("user with username %s not found", userName)
		}
		return model.User{}, fmt.Errorf("failed to get user by username %v", userName)
	}

	return user, nil
}

func UpdatePassword() error { return nil }

func UpdateUserName() error { return nil }

func UpdateEmail() error { return nil }

func UpdateName() error { return nil }

func DeleteAccount() error { return nil }

func CloseAccount() error { return nil }

func RestoreAccount() error { return nil }
