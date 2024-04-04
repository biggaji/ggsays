package routes

import (
	"github.com/biggaji/ggsays/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRoute := app.Group("users")

	userRoute.Post("/", handlers.HandleCreateUser)
	userRoute.Post("/authenticate", handlers.HandleUserAuthentication)
}
