package models

type EveningRequest struct {
	ID           uint    `gorm:"primarykey"`
	Expenses     float64 `json:"expenses"`
	Status       string  `json:"string"`
	Remaining    float64 `json:"remaining"`
	MorningID    uint    `gorm:"not null"`
}
