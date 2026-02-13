package models

import "github.com/google/uuid"

type Notification struct {
	BaseModel

	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	ActorID uuid.UUID `gorm:"type:uuid;not null;index;"`
	PostID *uuid.UUID `gorm:"type:uuid;index;"`

	Type string `gorm:"size:50;not null;index`

	Read bool `gorm:"default:false;index"`
}
 