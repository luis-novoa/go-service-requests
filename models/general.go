package models

import (
  "time"
  "github.com/jinzhu/gorm"
)

type General struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
}

