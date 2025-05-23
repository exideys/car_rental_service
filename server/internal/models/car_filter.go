package models

type CarFilter struct {
	MinPrice     uint32   `form:"price-from"`
	MaxPrice     uint32   `form:"price-to"`
	CarBrand     string   `form:"car_brand"`
	YearOfIssues uint16   `form:"year_of_issue"`
	CarClass     []string `form:"car-class" binding:"omitempty,dive,oneof=Premium Comfort Econom"`
}
