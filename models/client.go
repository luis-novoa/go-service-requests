package models

import "github.com/jinzhu/gorm"

type Client struct {
	General
	Name string `gorm:"size:75;unique;not null"`
	AuthToken string `gorm:"size:64;unique;not null"`
	ServiceRequests []ServiceRequest
}