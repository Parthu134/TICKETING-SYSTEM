package routes

import (
	"balance-sheet/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.LoginUser)
	app.Post("/morning", handlers.MorningRoutes)
	app.Post("/evening", handlers.EveningRoutes)
}
