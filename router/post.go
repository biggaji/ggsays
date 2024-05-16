package router

import (
	"os"

	"github.com/biggaji/ggsays/handler"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupPostRoutes(app *fiber.App) {
	postRoute := app.Group("posts")

	postRoute.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
	}))

	postRoute.Post("/", handler.HandleCreateNewPost)
	postRoute.Get("/:id", handler.HandleGetPostById)
	postRoute.Get("/", handler.HandleGetPosts)
}
