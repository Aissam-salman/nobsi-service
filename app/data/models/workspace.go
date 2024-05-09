package models

import (
	"time"

)


type Workspace struct {
    ID        uint       `gorm:"primaryKey" json:"id"`
    Name      string     `gorm:"not null" json:"name"`
    OwnerID   uint       `gorm:"not null" json:"ownerID"`
    Folders   []Folder   `json:"folders"`
    CreatedAt time.Time  `json:"createdAt"`
    UpdatedAt time.Time  `json:"updatedAt"`
}