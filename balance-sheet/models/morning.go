package models

import "time"

type MorningRequest struct {
	ID             uint      `gorm:"primarykey"`
	Date           time.Time `gorm:"unique;not null"`
	UserID         uint      `json:"user_id"`
	OpeningCash    float64   `json:"openingcash"`
	OpeningWallet  float64   `json:"openingwallet"`
	OpeningAccount float64   `json:"openingaccount"`
	TotalCash      float64   `json:"total_cash"`
}
