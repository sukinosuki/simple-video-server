package collection

import "time"

type AddCollection struct {
	VID uint `json:"vid" form:"vid" binding:"required"`
}

type UserVideoCollectionRes struct {
	VID       uint      `json:"vid"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
