package test_student

import "time"

type Student struct {
	ID   uint   `json:"id"`
	Name string `json:"name" gorm:"not null;size:12;unique;"`
	//Information Information `json:"information" gorm:"foreignKey:student_id"`
	Information Information `json:"information"`
	Book        []Book      `json:"book"`
	Language    []Language  `json:"language" gorm:"many2many:student_language"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// Book 一个student有多个book 一对多
type Book struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" gorm:"not null;uniqueIndex:idx_name_student;size:50;"` // gorm设置联合唯一索引(给两个或多个字段加上uniqueIndex:[同样的索引名]): uniqueIndex:idx_name_student
	StudentID uint      `json:"student_id" gorm:"not null;uniqueIndex:idx_name_student"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Information 一个学生有一个information 一对一
type Information struct {
	ID        uint      `json:"id"`
	Age       int       `json:"age" gorm:"not null;"`
	StudentID uint      `json:"student_id" gorm:"not null;unique;"` //为了限制一个student只能有一条information记录，这时给student_id(外键)加上唯一索引
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Language 每个student都可以有多个language
type Language struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" gorm:"unique;size:30;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Student   []Student `json:"student" gorm:"many2many:student_language"`
}

//type StudentLanguage struct {
//	ID         uint      `json:"id"`
//	CreatedAt  time.Time `json:"created_at"`
//	UpdatedAt  time.Time `json:"updated_at"`
//	StudentID  uint      `json:"student_id"`
//	LanguageID uint      `json:"language_id"`
//}

type StudentAdd struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type InformationAdd struct {
	Age       int  `json:"age" form:"age" binding:"required"`
	StudentID uint `json:"student_id" form:"student_id" binding:"required"`
}

type BookAdd struct {
	Name      string `json:"name" form:"book" binding:"required"`
	StudentID uint   `json:"student_id" form:"book" binding:"required"`
}

type LanguageAdd struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type BindStudentAndLanguage struct {
	StudentID uint `json:"student_id" form:"student_id" binding:"required"`

	LanguageID []uint `json:"language_id" form:"language_id" binding:"required"`
}
