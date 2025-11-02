package models

import "time"

type Order struct {
	OrderID    uint      `json:"order_id" gorm:"primaryKey"`
	ClientID   uint      `json:"client_id"`
	CarID      uint      `json:"car_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Status     string    `json:"status"`
	IsPaid     bool      `json:"is_paid"`
	TotalCost  float64   `json:"total_cost"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Order) TableName() string { return "orders" }
