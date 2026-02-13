package models

import "github.com/google/uuid"

type Repost struct {
	BaseModel

	UserId uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_post_repost"`
	PostId uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_user_post_repost"`
}
