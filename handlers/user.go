package handlers

import (
	"strings"

	"github.com/biggaji/ggsays/dtos"
	"github.com/biggaji/ggsays/helpers"
	"github.com/biggaji/ggsays/models"
	"github.com/biggaji/ggsays/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandleCreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Basic Validation
	if !strings.Contains(user.Email, "@") {
		return helpers.CreateErrorResponse(c, "a valid email is required", fiber.StatusBadRequest)
	}

	if len(user.UserName) < 3 {
		return helpers.CreateErrorResponse(c, "username is required and should be at least 3 characters long", fiber.StatusBadRequest)
	}

	if user.FirstName == "" {
		return helpers.CreateErrorResponse(c, "firstname is required", fiber.StatusBadRequest)
	}

	if user.LastName == "" {
		return helpers.CreateErrorResponse(c, "lastname is required", fiber.StatusBadRequest)
	}

	if len(user.Password) < 8 {
		return helpers.CreateErrorResponse(c, "password is required and it should not be less than 8 characters", fiber.StatusBadRequest)
	}

	userExist := repository.UserRecordExist(user.Email)

	if userExist {
		return helpers.CreateErrorResponse(c, "user account exist already", fiber.StatusBadRequest)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		panic(err)
	}

	user.ID = uuid.New()
	user.Password = string(hash)

	repository.InsertUserRecord(user)

	response := helpers.CreateUserResponse(user)

	return c.Status(fiber.StatusCreated).JSON(response)
}

func HandleUserAuthentication(c *fiber.Ctx) error {
	var payload dtos.UserAuthentication

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// Basic Validation
	if !strings.Contains(payload.Email, "@") {
		return helpers.CreateErrorResponse(c, "a valid email is required", fiber.StatusBadRequest)
	}

	if payload.Password == "" {
		return helpers.CreateErrorResponse(c, "password is required", fiber.StatusBadRequest)
	}

	userExist := repository.UserRecordExist(payload.Email)

	if !userExist {
		return helpers.CreateErrorResponse(c, "user account not found", fiber.StatusNotFound)
	}

	user, err := repository.GetUserByEmail(payload.Email)

	if err != nil {
		return helpers.CreateErrorResponse(c, "an error occured on our end", fiber.StatusInternalServerError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return helpers.CreateErrorResponse(c, "invalid login credentials", fiber.StatusBadRequest)
	}

	accessToken, err := helpers.GenerateAccessToken(dtos.UserJwtPayload{ID: user.ID})

	if err != nil {
		return helpers.CreateErrorResponse(c, "an error occured on our end", fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken": accessToken,
	})
}
