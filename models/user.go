package models

import "time"

// User User结构体默认的表名为`users`, 如果需要自定义表名, 可以让User实现TableName方法
type User struct {
	//gorm.Model
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted_at"`
	Nickname  string     `json:"nickname" gorm:"not null;size:12;type:string;comment:昵称"`
	Email     string     `json:"email" gorm:"uniqueIndex;not null;size:50;type:string;comment:邮箱;"`
	Password  string     `json:"-" gorm:"not null;size:255;type:string"`
	Enabled   bool       `json:"enabled" gorm:"not null;type:bool"`
	Cover     string     `json:"cover" gorm:"not null;size(255)"`
	//VideoCollections []Video `json:"video_collections" gorm:"many2many:user_collection_video;"`
}
