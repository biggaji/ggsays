package routes

import (
	"os"

	"github.com/biggaji/ggsays/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(app *fiber.App) {
	postRoute := app.Group("posts")

	postRoute.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
	}))

	postRoute.Post("/", handlers.HandleCreateNewPost)
	postRoute.Get("/:id", handlers.HandleGetPostById)
	postRoute.Get("/", handlers.HandleGetPosts)
}
