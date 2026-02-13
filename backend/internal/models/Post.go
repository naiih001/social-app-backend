package models

import "github.com/google/uuid"

type Post struct {
	BaseModel

	UserId uuid.UUID `gorm:"type:uuid;not null;index"`
	Content string `gorm;type:text;not null`

	ParentPostId *uuid.UUID `gorm:"type:uuid;index"`

	LikesCount int64 `gorm:"default:0"`
	RepostsCount int64 `gorm:"default:0"`
	CommentsCount int64 `gorm:"default:0"`
}

