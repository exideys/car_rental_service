package models

import (
	"time"
)

type Order struct {
	OrderID   uint      `gorm:"primaryKey" json:"id"`
	ClientID  uint      `json:"client_id"`
	CarID     uint      `json:"car_id"`
	TariffID  uint      `json:"tariff_id"`
	IsPaid    bool      `json:"is_paid"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	TotalCost float64   `json:"total_cost"`
	Status    string    `json:"status"`
}

func (Order) TableName() string { return "orders" }
