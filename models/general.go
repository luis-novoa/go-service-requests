package models

import "time"

type General struct {
  ID        int `gorm:"primary_key" json:"id,omitempty"`
  CreatedAt time.Time `json:"created_at,omitempty"`
  UpdatedAt time.Time `json:"updated_at,omitempty"`
}

