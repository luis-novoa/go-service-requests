package models

import "time"

type ServiceRequest struct {
	ID        int `gorm:"primary_key" json:"id,omitempty"`
  CreatedAt time.Time `json:"created_at,omitempty"`
  UpdatedAt time.Time `json:"updated_at,omitempty"`
	Status string `gorm:"not null" json:"status,omitempty"`
	Review int `json:"review,omitempty"`
	ClientID int `gorm:"not null" json:"client_id,omitempty"`
	Client User `gorm:"foreignKey:ClientID"`
	TechnicianID int `gorm:"not null" json:"technician_id,omitempty"`
	Technician User `gorm:"foreignKey:TechnicianID"`
}