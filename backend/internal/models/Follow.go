package models

import "github.com/google/uuid"

type Follow struct {
	BaseModel

	FollowerID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_follow_unique"`
	FolloweeID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:idx_follow_unique"`

	// Prevent duplitates
	_ struct{} `gorm"uniqueIndex:idx_follow_unique,priority:1"`
	_ struct{} `gorm:"uniqueIndex:idx_follow_unique,priority:2"` 
}