package models

import "gorm.io/gorm"

type Pqr struct {
	gorm.Model
	Type        int    `json:"type"`
	Email       string `json:"email"`
	Description string `json:"description"`
	FkCustomer  int    `json:"fk_customer"`
}
