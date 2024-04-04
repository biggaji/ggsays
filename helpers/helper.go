package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/biggaji/ggsays/dtos"
	"github.com/biggaji/ggsays/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateUserResponse(userModel models.User) dtos.User {
	return dtos.User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName, Email: userModel.Email, UserName: userModel.UserName, Password: userModel.Password}
}

func CreateErrorResponse(c *fiber.Ctx, message string, status int) error {
	return c.Status(status).JSON(fiber.Map{"error": errors.New(message).Error()})
}

func GenerateAccessToken(payload dtos.UserJwtPayload) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	accessToken, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return accessToken, nil
}

func CreatePostResponse(postModel models.Post, includeUser bool) dtos.Post {
	post := dtos.Post{
		ID:        postModel.ID,
		LikeCount: postModel.LikeCount,
		Content:   postModel.Content,
		UserId:    postModel.UserId,
	}

	if includeUser {
		post.User = &postModel.User
	} else {
		post.User = nil
	}

	return post
}

func CreatePostsResponse(postModel []models.Post, includeUser bool) []dtos.Post {
	var posts []dtos.Post

	for _, post := range postModel {
		data := dtos.Post{
			ID:        post.ID,
			LikeCount: post.LikeCount,
			Content:   post.Content,
			UserId:    post.UserId,
		}

		if includeUser {
			data.User = &post.User
		} else {
			data.User = nil
		}
		posts = append(posts, data)
	}

	return posts
}

func ExtractUserIdFromJwtClaims(c *fiber.Ctx) (uuid.UUID, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userId, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, errors.New("user ID not found in JWT claims")
	}

	return uuid.MustParse(userId), nil
}
