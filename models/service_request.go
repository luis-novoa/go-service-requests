package models

import "github.com/jinzhu/gorm"

type ServiceRequest struct {
	General
	Status string `gorm:"not null"`
	Review int
	UserID
	User User
	TechnicianID
	Technician Technician
}