package models

import (
	"time"

)

type Folder struct {
    ID          uint       `gorm:"primaryKey" json:"id"`
    Name        string     `gorm:"not null" json:"name"`
    WorkspaceID uint       `gorm:"not null" json:"workspaceID"`
    Files       []File     `json:"files"`
    CreatedAt   time.Time  `json:"createdAt"`
    UpdatedAt   time.Time  `json:"updatedAt"`
}