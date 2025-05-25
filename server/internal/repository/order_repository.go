package repository

import (
	"github.com/exideys/car_rental_service/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	FindByEmail(email string) (*models.Client, error)
	GetAllOrders(email string) ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindByEmail(email string) (*models.Client, error) {
	var client models.Client
	if err := r.db.Where("email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *orderRepository) GetAllOrders(email string) ([]models.Order, error) {
	client, err := r.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	var orders []models.Order
	if err := r.db.
		Where("client_id = ?", client.ClientID).
		Find(&orders).
		Error; err != nil {
		return nil, err
	}
	return orders, nil
}
