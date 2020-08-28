package models

type ServiceRequest struct {
	General
	Status string `gorm:"not null;DEFAULT:Requested" json:"status,omitempty"`
	Review int `json:"review,omitempty"`
	ClientID int `gorm:"not null" json:"client_id,omitempty"`
	Client User `gorm:"foreignKey:ClientID"`
	TechnicianID int `gorm:"not null" json:"technician_id,omitempty"`
	Technician User `gorm:"foreignKey:TechnicianID"`
}