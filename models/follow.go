package models

import "time"

type Follow struct {
	UID       uint      `json:"uid" gorm:"primaryKey"`
	TargetUID uint      `json:"target_uid" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;"`
}
