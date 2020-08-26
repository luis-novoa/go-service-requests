package models

import "github.com/jinzhu/gorm"

type Technician struct {
	Client
	ServiceRequests []ServiceRequest
}