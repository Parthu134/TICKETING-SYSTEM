package handlers

import (
	"balance-sheet/config"
	"balance-sheet/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *fiber.Ctx) error {
	var input AuthInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user_id": user.ID,
	})
}
