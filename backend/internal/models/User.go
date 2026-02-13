package models

type User struct {
	BaseModel

	Username string `gorm:"uniqueIndex;size:30;not null"`
	Email    string `gorm:"uniqueIndex;size:100;not null"`
	Password string `gorm:"not null"`

	Name string `gorm:"size:100"`
	Bio  string `gorm:"size:280"`

	AvatarURL string `gorm:"size:255"`

	FollowersCount int64 `gorm:"default:0"`
	FollowingCount int64 `gorm:"default:0"`
	PostsCount     int64 `gorm:"default:0"`
}
