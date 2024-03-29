package form

import (
	"camarinb2096/form_example/internal/app/complaint"
	"camarinb2096/form_example/internal/app/customer"
	"errors"

	"gorm.io/gorm"
)

type Services interface {
	CreatePqr(name, email, city, complaint string) error
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Services {
	return &service{
		db: db,
	}
}
func (s *service) CreatePqr(name, email, city, comp string) error {
	//TODO: Add validation for email and city, add validation for complaint, change form for complaint, add http responses for errors and success

	customerService := customer.NewCustomerService(s.db)

	fkCustomerId, err := customerService.GetCustomerId(email)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Si hay un error que no es ErrRecordNotFound, devuélvelo
			return err
		}
		// Si el error es ErrRecordNotFound, crea un nuevo cliente
		fkCustomerId, err = customerService.CreateCustomer(name, email, city)
		if err != nil {
			return err
		}
	} else if fkCustomerId != 0 {
		return errors.New("el usuario ya ha procesado una PQR anteriormente")
	}

	complaintService := complaint.NewComplaintService(s.db)

	if err := complaintService.CreateComplaint("Pendiente", "PQR", comp, int(fkCustomerId)); err != nil {
		return err
	}

	return nil
}
