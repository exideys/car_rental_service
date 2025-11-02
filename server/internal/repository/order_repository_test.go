package repository

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/exideys/car_rental_service/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Add AnyTime type for mocking time fields
type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestOrderRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := NewOrderRepository(gormDB)

	testOrder := &models.Order{
		OrderID:    1,
		ClientID:   1,
		CarID:      1,
		StartDate:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
		Status:     "pending",
		IsPaid:     false,
		TotalCost:  100.0,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "orders"`).
		WithArgs(testOrder.ClientID, testOrder.CarID, testOrder.StartDate, testOrder.EndDate, testOrder.Status, testOrder.IsPaid, testOrder.TotalCost, AnyTime{}, AnyTime{}, testOrder.OrderID).
		WillReturnRows(sqlmock.NewRows([]string{"order_id"}).AddRow(1))
	mock.ExpectCommit()

	err = repo.Create(testOrder)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderRepository_FindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := NewOrderRepository(gormDB)

	email := "test@example.com"
	client := &models.Client{
		ClientID:  1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     email,
	}

	rows := sqlmock.NewRows([]string{"client_id", "first_name", "last_name", "email"}).
		AddRow(client.ClientID, client.FirstName, client.LastName, client.Email)

	mock.ExpectQuery(`SELECT \* FROM "client" WHERE email = \$1 ORDER BY "client"."client_id" LIMIT \$2`).
		WithArgs(email, 1).
		WillReturnRows(rows)

	foundClient, err := repo.FindByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, client, foundClient)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderRepository_GetAllOrders(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := NewOrderRepository(gormDB)

	email := "test@example.com"
	client := &models.Client{
		ClientID: 1,
		Email:    email,
	}
	orders := []models.Order{
		{OrderID: 1, ClientID: 1, CarID: 1},
		{OrderID: 2, ClientID: 1, CarID: 2},
	}

	clientRows := sqlmock.NewRows([]string{"client_id", "email"}).
		AddRow(client.ClientID, client.Email)
	mock.ExpectQuery(`SELECT \* FROM "client" WHERE email = \$1 ORDER BY "client"."client_id" LIMIT \$2`).
		WithArgs(email, 1).
		WillReturnRows(clientRows)

	orderRows := sqlmock.NewRows([]string{"order_id", "client_id", "car_id"}).
		AddRow(orders[0].OrderID, orders[0].ClientID, orders[0].CarID).
		AddRow(orders[1].OrderID, orders[1].ClientID, orders[1].CarID)
	mock.ExpectQuery(`SELECT \* FROM "orders" WHERE client_id = \$1`).
		WithArgs(client.ClientID).
		WillReturnRows(orderRows)

	foundOrders, err := repo.GetAllOrders(email)
	assert.NoError(t, err)
	assert.Len(t, foundOrders, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderRepository_GetAllOrders_ClientNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := NewOrderRepository(gormDB)

	email := "nonexistent@example.com"

	mock.ExpectQuery(`SELECT \* FROM "client" WHERE email = \$1 ORDER BY "client"."client_id" LIMIT \$2`).
		WithArgs(email, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	_, err = repo.GetAllOrders(email)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
