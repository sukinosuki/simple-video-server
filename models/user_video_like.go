package models

import "time"

type UserVideoLike struct {
	VID       uint `json:"vid" gorm:"primaryKey;column:vid;"`
	UID       uint `json:"uid" gorm:"not null;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
