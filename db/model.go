package db

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel defines ID and audit timestamps for all models.
type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
