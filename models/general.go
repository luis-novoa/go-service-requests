package models

import (
  "time"
  "encoding/json"
  "github.com/jinzhu/gorm"
)

type General struct {
  ID        uint `gorm:"primary_key" json:"id,omitempty"`
  CreatedAt time.Time `json:"created_at,omitempty"`
  UpdatedAt time.Time `json:"updated_at,omitempty"`
}

