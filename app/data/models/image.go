package models

import (
	"time"
)


type Image struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"not null" json:"name"`
    Path      string    `gorm:"not null" json:"path"`
    FileID uint     `gorm:"not null" json:"documentID"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}