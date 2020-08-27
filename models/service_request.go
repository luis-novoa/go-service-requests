package models

import "github.com/jinzhu/gorm"

type ServiceRequest struct {
	General
	Status string `gorm:"not null;DEFAULT:Requested"`
	Review int
	UserID `gorm:"not null"`
	User User
	TechnicianID `gorm:"not null"`
	Technician Technician
}