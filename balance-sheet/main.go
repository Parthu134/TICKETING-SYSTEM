package main

import (
	"balance-sheet/config"
	"balance-sheet/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()
	routes.Setup(app)
	app.Listen(":3000")
}
