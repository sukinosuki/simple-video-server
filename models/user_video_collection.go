package models

import "time"

type UserVideoCollection struct {
	VID       uint      `json:"vid" gorm:"column:vid;not null;primaryKey"`
	UID       uint      `json:"uid" gorm:"not null;primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;"`
}
