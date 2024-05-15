package route

import (
	"github.com/biggaji/ggsays/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRoute := app.Group("users")

	userRoute.Post("/", handler.HandleCreateUser)
	userRoute.Post("/authenticate", handler.HandleUserAuthentication)
}
