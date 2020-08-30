package models

import "time"

type User struct {
	ID        int `gorm:"primary_key" json:"id,omitempty"`
  CreatedAt time.Time `json:"created_at,omitempty"`
  UpdatedAt time.Time `json:"updated_at,omitempty"`
	Name string `gorm:"size:75;unique;not null" json:"name,omitempty"`
	AuthToken string `gorm:"size:64;unique;not null" json:"auth_token,omitempty"`
	Technician bool `gorm:"not null" json:"technician,omitempty"`
	ServiceRequests []ServiceRequest
}