package handlers

import (
	"balance-sheet/config"
	"balance-sheet/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MorningInput struct {
	UserID         uint    `json:"user_id"`
	OpeningCash    float64 `json:"openingcash"`
	OpeningAccount float64 `json:"openingaccount"`
	OpeningWallet  float64 `json:"openingWallet"`
}

func MorningRoutes(c *fiber.Ctx) error {
	var input MorningInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	entryDate := time.Now().Truncate(24 * time.Hour)
	var existing models.MorningRequest
	if err := config.DB.Where("date=?", entryDate).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "morning entry already exists for this date",
		})
	}
	totalCash := input.OpeningAccount + input.OpeningCash + input.OpeningWallet
	morning := models.MorningRequest{
		Date:           entryDate,
		UserID:         input.UserID,
		OpeningCash:    input.OpeningCash,
		OpeningWallet:  input.OpeningWallet,
		OpeningAccount: input.OpeningAccount,
		TotalCash:      totalCash,
	}

	if err := config.DB.Create(&morning).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "failed to save morning data",
		})
	}
	return c.JSON(fiber.Map{
		"message":         "morning balances recorder",
		"date":            morning.Date,
		"opening_balance": morning.TotalCash,
	})
}
