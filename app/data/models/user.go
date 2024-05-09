package models

import (
	"time"
)

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `gorm:"unique;not null" json:"username"`
    Email     string    `gorm:"unique;not null" json:"email"`
    Password  string    `gorm:"not null" json:"-"`
    WorkspaceID uint    `gorm:"not null" json:"workspaceID"`
    Workspace Workspace `json:"workspace"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
