package models

import "github.com/google/uuid"

type Comment struct {
	BaseModel

	PostID uuid.UUID `gorm:"type:uuid;not null;index;"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index;"`

	ParentCommentID *uuid.UUID `gorm:"type:uuid;index;"`

	Content string `gorm:"type:text;not null;"`
}
