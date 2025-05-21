package models

import "time"

type ClientProfile struct {
	ClientID           int64      `json:"client_id"`
	FirstName          string     `json:"first_name"`
	LastName           string     `json:"last_name"`
	Email              string     `json:"email"`
	BirthDate          time.Time  `json:"birth_date"`
	Age                int        `json:"age"`
	IsBlocked          bool       `json:"is_blocked"`
	IsVIP              bool       `json:"is_vip"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	TotalOrders        int64      `json:"total_orders"`
	TotalPaid          float64    `json:"total_paid"`
	TotalEvaluations   int64      `json:"total_evaluations"`
	AvgRating          float64    `json:"avg_rating"`
	LastEvaluationDate *time.Time `json:"last_evaluation_date,omitempty"`
	LateReturnCount    int64      `json:"late_return_count"`
}
