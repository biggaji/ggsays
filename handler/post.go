package handler

import (
	"github.com/biggaji/ggsays/helper"
	"github.com/biggaji/ggsays/model"
	"github.com/biggaji/ggsays/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleCreateNewPost(c *fiber.Ctx) error {
	var post model.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if len(post.Content) < 25 {
		return helper.CreateErrorResponse(c, "post content must be at least 25 characters long", fiber.StatusBadRequest)
	}

	userId, err := helper.ExtractUserIdFromJwtClaims(c)
	if err != nil {
		return helper.CreateErrorResponse(c, "failed to extract user ID from JWT claims", fiber.StatusInternalServerError)
	}

	// Check if user has already posted that before
	contentExist := repository.PostContentExistForUser(userId, post.Content)

	if contentExist {
		return helper.CreateErrorResponse(c, "oops you've already posted that, try posting an new content", fiber.StatusConflict)
	}

	post.ID = uuid.New()
	post.UserId = userId

	repository.InsertNewPostRecord(post)

	response := helper.CreatePostResponse(post, false)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "post": response})
}

func HandleGetPostById(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return helper.CreateErrorResponse(c, "post id is required", fiber.StatusBadRequest)
	}

	postId := uuid.MustParse(id)

	post, err := repository.GetPostById(postId, true)

	if err != nil {
		return helper.CreateErrorResponse(c, "post not found", fiber.StatusNotFound)
	}

	response := helper.CreatePostResponse(post, false)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "post": response})
}

func HandleGetPosts(c *fiber.Ctx) error {

	posts := repository.GetAllPosts(true)

	response := helper.CreatePostsResponse(posts, false)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "post": response})
}

func HandlePostUpvote(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return helper.CreateErrorResponse(c, "post id is required", fiber.StatusBadRequest)
	}

	postId := uuid.MustParse(id)

	post, err := repository.GetPostById(postId, true)

	if err != nil {
		return helper.CreateErrorResponse(c, "post not found", fiber.StatusNotFound)
	}

	// Increment like by one
	post.LikeCount += 1

	// Also check if user has liked the post before, and either unlike the post or return a response indicating user can only like post once.

	return c.SendStatus(fiber.StatusOK)
}
