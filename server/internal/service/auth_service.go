package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/exideys/car_rental_service/internal/models"
	"regexp"
	"strings"
	"time"

	"github.com/exideys/car_rental_service/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	SignUp(ctx context.Context, firstName, lastName, email, telephone, password, passwordConfirm, birthDate string) error
	Login(ctx context.Context, email, password string) (*models.Client, error)
	GetByEmail(email string) (*models.Client, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) SignUp(ctx context.Context, first_name, last_name, email, telephone, password, password_confirm, birth_date string) error {
	first_name = strings.TrimSpace(first_name)
	last_name = strings.TrimSpace(last_name)
	email = strings.TrimSpace(email)
	telephone = strings.TrimSpace(telephone)
	password = strings.TrimSpace(password)
	password_confirm = strings.TrimSpace(password_confirm)
	birth_date = strings.TrimSpace(birth_date)

	// Empty checks
	if first_name == "" {
		return errors.New("first name is required")
	}
	if last_name == "" {
		return errors.New("last name is required")
	}
	if email == "" {
		return errors.New("email is required")
	}
	if telephone == "" {
		return errors.New("telephone is required")
	}
	if password == "" {
		return errors.New("password is required")
	}
	if password_confirm == "" {
		return errors.New("confirm password is required")
	}
	if birth_date == "" {
		return errors.New("birth date is required")
	}

	// Password confirmation
	if password != password_confirm {
		return errors.New("password and confirmation do not match")
	}

	// Email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	// Password length check (min 6 chars)
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// Telephone validation (basic: only digits, min 10)
	telephoneRegex := regexp.MustCompile(`^\+?\d{10,15}$`)
	if !telephoneRegex.MatchString(telephone) {
		return errors.New("invalid telephone number format")
	}

	// Birthdate validation (YYYY-MM-DD)
	bd, err := time.Parse("2006-01-02", birth_date)
	if err != nil {
		return errors.New("birth date must be in YYYY-MM-DD format")
	}
	if _, err := s.repo.FindByEmail(ctx, email); err == nil {
		return errors.New("email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	return s.repo.SignUp(ctx, first_name, last_name, email, telephone, string(hashed), bd.Format("2006-01-02"))

}

func (s *authService) Login(ctx context.Context, email, password string) (*models.Client, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	client, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("incorrect email or password")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password.Hash), []byte(password)); err != nil {
		return nil, errors.New("incorrect email or password")
	}

	if client.IsBlocked {
		return nil, errors.New("account is blocked")
	}

	return client, nil
}

func (s *authService) GetByEmail(email string) (*models.Client, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return nil, errors.New("email is required")
	}
	return s.repo.FindByEmail(context.Background(), email)
}
