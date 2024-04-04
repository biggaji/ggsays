package main

import (
	"fmt"
	"log"
	"os"

	"github.com/biggaji/ggsays/database"
	"github.com/biggaji/ggsays/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env")
	}

	app := fiber.New()

	database.Connect()

	routes.SetupUserRoutes(app)
	routes.SetupPostRoutes(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	app.Listen(fmt.Sprintf(":%v", port))
}
