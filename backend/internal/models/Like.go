package models

import "github.com/google/uuid"

type Like struct {
	BaseModel

	UserId uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_post_like"`
	PostId uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_post_like"`
}
