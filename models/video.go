package models

import "time"

type Video struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted_at"`
	Cover     string     `json:"cover" gorm:"not null;type:string;size:255;"`
	Snapshot  string     `json:"snapshot" gorm:"not null;type:string;size:255;"`
	Title     string     `json:"title" gorm:"not null;type:string;size:50;"`
	Locked    bool       `json:"locked" gorm:"not null;type:bool;comment:0不锁定1锁定"`
	Status    int        `json:"status" gorm:"not null;size:2;comment:0审核中1审核不通过2审核通过"`
	Url       string     `json:"url" gorm:"not null;type:string;size:255;"`
	Uid       uint       `json:"uid" gorm:"not null;index;"`
	//User      *User      `json:"user" gorm:"embedded;embeddedPrefix:user_"`
	// TODO: video 分类、标签
}
