package models

import (
	"time"

)

type File struct {
    ID          uint       `gorm:"primaryKey" json:"id"`
    Name        string     `gorm:"not null" json:"name"`
    Content     string    `gorm:"type:text;not null" json:"content"`
    Size        int64      `gorm:"not null" json:"size"`
    Path        string     `gorm:"not null" json:"path"` // Chemin d'accès au fichier
	Images      []Image   `json:"images"`
    FolderID    uint       `gorm:"not null" json:"folderID"`
    CreatedAt   time.Time  `json:"createdAt"`
    UpdatedAt   time.Time  `json:"updatedAt"`
}
