package handlers

import (
	"balance-sheet/config"
	"balance-sheet/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type EveningInput struct {
	Expenses float64 `json:"expenses"`
}

func EveningRoutes(c *fiber.Ctx) error {
	var input EveningInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	entryDate := time.Now().Truncate(24 * time.Hour)
	startDay := entryDate
	endday := entryDate.AddDate(0, 0, 1)

	var morning models.MorningRequest
	if err := config.DB.Where("date>=? AND date<?", startDay, endday).First(&morning).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "no morning entry is found today",
		})
	}

	openingBalance := morning.TotalCash
	remaining := openingBalance - input.Expenses
	status := ""
	if remaining > openingBalance {
		status = "Profit"
	} else if remaining < openingBalance {
		status = "Loss"
	} else {
		status = "no profit/loss"
	}
	evening := models.EveningRequest{
		MorningID: morning.ID,
		Expenses:  input.Expenses,
		Status:    status,
		Remaining: remaining,
	}
	if err := config.DB.Create(&evening).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save evening details",
		})
	}
	return c.JSON(fiber.Map{
		"date":            entryDate.Format("2006-01-02"),
		"opening_balance": openingBalance,
		"total_expenses":  evening.Expenses,
		"remaining":       remaining,
		"status":          status,
	})
}
