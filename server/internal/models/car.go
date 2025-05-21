package models

type Status string

const (
	StatusAvailable   Status = "Available"
	StatusRented      Status = "Not available"
	StatusMaintenance Status = "Under repair"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusAvailable, StatusRented, StatusMaintenance:
		return true
	}
	return false
}

func (s Status) String() string {
	return string(s)
}

type Car struct {
	CarID        uint   `gorm:"primaryKey;column:car_id" json:"car_id"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Color        string `json:"color"`
	PlateNumber  string `json:"plate_number"`
	YearOfIssue  uint   `json:"year_of_issue"`
	CarClass     string `json:"car_class"`
	InsuranceNum string `gorm:"type:varchar(100)" json:"insurance_num"`
	Status       Status `gorm:"type:varchar(20);default:'Вільний'" json:"status"`
	Daily_price  uint   `json:"daily_price"`
	ImagePath    string `json:"image_path"`
}

func (Car) TableName() string {
	return "cars"
}
