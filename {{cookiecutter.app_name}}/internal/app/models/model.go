package models

import (
	"time"
)

// Timestamps Timestamp columns
type Timestamps struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;not null;default:NOW()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime;not null;default:NOW()"`
}

// BaseTimestamps Timestamp columns with deleted_at
type BaseTimestamps struct {
	Timestamps
	DeletedAt *time.Time `json:"-" sql:"index" gorm:"type:datetime;"`
}
