package models

type Status string

const (
	StatusAvailable   Status = "Вільний"
	StatusRented      Status = "Недоступний"
	StatusMaintenance Status = "На ремонті"
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

type TransportType string

const (
	TransportSedan     TransportType = "Седан"
	TransportSUV       TransportType = "Позашляховик"
	TransportHatchback TransportType = "Хетчбек"
)

func (t TransportType) IsValid() bool {
	switch t {
	case TransportSedan, TransportSUV, TransportHatchback:
		return true
	}
	return false
}

func (t TransportType) String() string {
	return string(t)
}

type Transmission string

const (
	TransmissionAutomatic Transmission = "Автомат"
	TransmissionManual    Transmission = "Механічна"
)

func (tr Transmission) IsValid() bool {
	switch tr {
	case TransmissionAutomatic, TransmissionManual:
		return true
	}
	return false
}

func (tr Transmission) String() string {
	return string(tr)
}

type CarClass string

const (
	CarClassPremium   CarClass = "Преміум"
	CarClassCrossover CarClass = "Кросовер"
	CarClassEconomy   CarClass = "Економ"
	CarClassCity      CarClass = "Місто"
)

func (cc CarClass) IsValid() bool {
	switch cc {
	case CarClassPremium, CarClassCrossover, CarClassEconomy, CarClassCity:
		return true
	}
	return false
}

func (cc CarClass) String() string {
	return string(cc)
}

type Car struct {
	CarID        uint     `gorm:"primaryKey;autoIncrement;column:car_id"        json:"car_id"`
	Brand        string   `gorm:"type:varchar(255);not null"                    json:"brand"`
	Model        string   `gorm:"type:varchar(255);not null"                    json:"model"`
	Color        string   `gorm:"type:varchar(50)"                              json:"color"`
	PlateNumber  string   `gorm:"size:191;uniqueIndex;column:plate_number;not null" json:"plate_number"`
	YearOfIssue  uint     `gorm:"column:year_of_issue"                          json:"year_of_issue"`
	CarClass     CarClass `gorm:"type:enum('Преміум','Кросовер','Економ','Місто');column:car_class" json:"car_class"`
	InsuranceNum string   `gorm:"type:varchar(100);column:insurance_num"        json:"insurance_num"`
	Status       Status   `gorm:"type:enum('Вільний','Недоступний','На ремонті');default:'Вільний';column:status" json:"status"`
	DailyPrice   float64  `gorm:"column:daily_price"                            json:"daily_price"`
}

func (Car) TableName() string {
	return "cars"
}
