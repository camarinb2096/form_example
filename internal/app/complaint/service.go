package complaint

import (
	"gorm.io/gorm"
)

type Services interface {
	CreateComplaint(status, CompType, complaint string, fkCustomer int) error
}

type services struct {
	db *gorm.DB
}

func NewComplaintService(db *gorm.DB) Services {
	return &services{
		db: db,
	}
}

//TODO: Add validation for complaint, add http responses for errors and success, add tests

func (s *services) CreateComplaint(status, CompType, complaint string, fkCustomer int) error {

	comp := Complaint{
		Status:     status,
		CompType:   CompType,
		Complaint:  complaint,
		FkCustomer: fkCustomer,
	}

	if err := s.db.Create(&comp).Error; err != nil {
		return err
	}

	return nil
}
