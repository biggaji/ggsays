package helper

import (
	"errors"
	"os"
	"time"

	"github.com/biggaji/ggsays/dto"
	"github.com/biggaji/ggsays/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateUserResponse(userModel model.User) dto.User {
	return dto.User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName, Email: userModel.Email, UserName: userModel.UserName, Password: userModel.Password}
}

func CreateErrorResponse(c *fiber.Ctx, message string, status int) error {
	return c.Status(status).JSON(fiber.Map{"error": errors.New(message).Error()})
}

func GenerateAccessToken(payload dto.UserJwtPayload) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": payload.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	accessToken, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", ErrCantGenerateAccessToken
	}

	return accessToken, nil
}

func CreatePostResponse(postModel model.Post, includeUser bool) dto.Post {
	post := dto.Post{
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

func CreatePostsResponse(postModel []model.Post, includeUser bool) []dto.Post {
	var posts []dto.Post

	for _, post := range postModel {
		data := dto.Post{
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

	// Handle type assertion failures
	// The single return value form of a type assertion will panic on an incorrect type. Therefore, always use the "comma ok" idiom.

	// t, ok := i.(string)
	// if !ok {
	//   // handle the error gracefully
	// }

	userId, ok := claims["sub"].(string)
	if !ok {
		return uuid.UUID{}, errors.New("user ID not found in JWT claims")
	}

	return uuid.MustParse(userId), nil
}
