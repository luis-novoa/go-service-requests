package models

type User struct {
	General
	Name string `gorm:"size:75;unique;not null" json:"name,omitempty"`
	AuthToken string `gorm:"size:64;unique;not null" json:"auth_token,omitempty"`
	Technician bool `gorm:"DEFAULT:false;not null" json:"technician,omitempty"`
	ServiceRequests []ServiceRequest
}