package models

type Car_filter struct {
	Min_price      uint32   `form:"price-from"`
	Max_price      uint32   `form:"price-to"`
	Car_brand      string   `form:"car-brand"`
	Year_of_issues uint16   `form:"Year_of_issues"`
	Car_class      []string `form:"car-class" binding:"omitempty,dive,oneof=Premium Comfort Econom"`
}
