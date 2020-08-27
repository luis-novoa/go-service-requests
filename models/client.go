package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Client struct {
	General
	Name string `gorm:"size:75;unique;not null" json:"name,omitempty"`
	AuthToken string `gorm:"size:64;unique;not null" json:"auth_token,omitempty"`
	ServiceRequests []ServiceRequest
}