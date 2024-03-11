package complaint

import "gorm.io/gorm"

type Complaint struct {
	gorm.Model
	Status     string `json:"status"`
	CompType   string `json:"type"`
	Complaint  string `json:"complaint"`
	FkCustomer int    `json:"fk_customer" gorm:"column:fk_customer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
