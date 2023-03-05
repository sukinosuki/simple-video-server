package test_student

import (
	"simple-video-server/pkg/global"
	"testing"
)

func TestDao_AddStudent(t *testing.T) {
	student := Student{
		Name: "hanami",
	}

	//err := Dao.AddStudent(&student)
	err := global.MysqlDB.Model(&Student{}).Create(student).Error

	if err != nil {
		t.Error(err)
	}

	t.Log("add student success, id: ", student.ID)
}
