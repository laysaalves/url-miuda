package entity

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key"`
	LinkLong string
	LinkMiudo string
	CreatedAt time.Time
}