package models

import (
	"time"

	"github.com/google/uuid"
)

type TimelineEntry struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserId uuid.UUID `gorm:"type:uuid;not null;index"`
	PostId uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time `gorm:"not null;index"`
}