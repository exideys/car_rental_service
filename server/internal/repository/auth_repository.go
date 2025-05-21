package repository

import (
	"context"
	"github.com/exideys/car_rental_service/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type AuthRepository interface {
	SignUp(ctx context.Context, first_name, last_name, email, telephone, hashed_password, birth_date string) error
	FindByEmail(ctx context.Context, email string) (*models.Client, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) SignUp(ctx context.Context, first_name, last_name, email, telephone, hashed_password, birth_date string) error {
	bd, err := time.Parse("2006-01-02", birth_date)
	if err != nil {
		return err
	}

	client := models.Client{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  models.Password{Hash: hashed_password},
		Telephone: telephone,
		BirthDate: bd,
		IsBlocked: false,
		IsVIP:     false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return r.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&client).Error
}

// FindByEmail looks up a client by e‑mail address.
func (r *authRepository) FindByEmail(ctx context.Context, email string) (*models.Client, error) {
	var client models.Client
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}
