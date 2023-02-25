package user

import "simple-video-server/models"

// User User结构体默认的表名为`users`, 如果需要自定义表名, 可以让User实现TableName方法
//type User struct {
//	//gorm.Model
//	ID        uint       `json:"id"`
//	CreatedAt time.Time  `json:"created_at" gorm:"not null"`
//	UpdatedAt time.Time  `json:"updated_at" gorm:"not null"`
//	DeletedAt *time.Time `json:"deleted_at"`
//	Nickname  string     `json:"nickname" gorm:"not null;size:12;type:string;comment:昵称"`
//	Email     string     `json:"email" gorm:"uniqueIndex;not null;size:50;type:string;comment:邮箱;"`
//	Password  string     `json:"-" gorm:"not null;size:255;type:string"`
//	Enabled   bool       `json:"enabled" gorm:"not null;type:bool"`
//}

//// TableName 自定义表名
//func (u *User) TableName() string {
//	return "user"
//}

type UserRegister struct {
	Email    string `json:"email" form:"email" binding:"required,max=50,min=6,email" label:"邮箱"`
	Password string `json:"password" form:"password" binding:"required,max=12,min=6" label:"密码"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginRes struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

//func init() {
//	err := global.MysqlDB.AutoMigrate(
//		&User{},
//	)
//
//	if err != nil {
//		panic(err)
//	}
//}
