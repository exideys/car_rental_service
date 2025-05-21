package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Client struct {
	ClientID  int64     `json:"client_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  Password  `json:"-" gorm:"column:password_hash;type:varchar(60)"`
	BirthDate time.Time `json:"birth_date"`
	Telephone string    `json:"telephone"`
	IsBlocked bool      `json:"is_blocked"`
	IsVIP     bool      `json:"is_vip" gorm:"column:is_vip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Client) TableName() string {
	return "client"
}

type Password struct {
	Plaintext *string
	Hash      string
}

func (p Password) Value() (driver.Value, error) {
	return p.Hash, nil
}

func (p *Password) Scan(v any) error {
	switch val := v.(type) {
	case []byte:
		p.Hash = string(val)
	case string:
		p.Hash = val
	case nil:
		p.Hash = ""
	default:
		return fmt.Errorf("cannot scan %T into Password", v)
	}
	return nil
}
