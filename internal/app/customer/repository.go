package customer

import (
	"camarinb2096/form_example/pkg/logger"

	"gorm.io/gorm"
)

type (
	Repository interface {
		CreateCustomer(name, email, city string) (int, error)
		GetCustomerId(email string) (int, error)
	}

	repo struct {
		db     *gorm.DB
		logger *logger.Logger
	}
)

func NewCustomerRepository(db *gorm.DB, logger *logger.Logger) Repository {
	return &repo{
		db:     db,
		logger: logger,
	}
}

func (r *repo) CreateCustomer(name, email, city string) (int, error) {
	customer := Customer{
		Name:  name,
		Email: email,
		City:  city,
	}

	if err := r.db.Create(&customer).Error; err != nil {
		return 0, err
	}

	return int(customer.ID), nil
}

func (r *repo) GetCustomerId(email string) (int, error) {
	var customer Customer

	r.logger.Info("Consultando usuario por email...")

	if err := r.db.Where("email = ?", email).First(&customer).Error; err != nil {
		return 0, err
	}

	return int(customer.ID), nil
}
