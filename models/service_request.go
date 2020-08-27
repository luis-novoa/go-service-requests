package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type ServiceRequest struct {
	General
	Status string `gorm:"not null;DEFAULT:Requested" json:"status,omitempty"`
	Review int `json:"review,omitempty"`
	UserID `gorm:"not null" json:"user_id,omitempty"`
	User User
	TechnicianID `gorm:"not null" json:"technician_id,omitempty"`
	Technician Technician
}