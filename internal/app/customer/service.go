package customer

import "gorm.io/gorm"

type Services interface {
	CreateCustomer(name, email, city string) (uint, error)
}

type services struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) Services {
	return &services{
		db: db,
	}
}

//TODO: Add validation for email and city, add http responses for errors and success, add tests

func (s *services) CreateCustomer(name, email, city string) (uint, error) {

	customer := Customer{
		Name:  name,
		Email: email,
		City:  city,
	}

	if err := s.db.Create(&customer).Error; err != nil {
		return 0, err
	}

	return customer.ID, nil
}
