package test_student

import (
	"gorm.io/gorm"
	"simple-video-server/pkg/global"
)

type dao struct {
}

var Dao = &dao{}

func (d *dao) AddStudent(student *Student) error {

	err := global.MysqlDB.Model(&Student{}).Create(student).Error

	return err
}

func (d *dao) AddInformation(information *Information) error {
	err := global.MysqlDB.Model(&Information{}).Create(information).Error

	return err
}

//func (d *dao) FindStudentById(id uint) (Information, error) {
//	var information Information
//	err := global.MysqlDB.Preload("Student").Where("student_id = ?", id).First(&information).Error
//
//	return information, err
//}

func (d *dao) FindStudentById(id uint) (Student, error) {
	var student Student
	//err := global.MysqlDB.Preload("book").Preload("Information").Where("id = ?", id).First(&student).Error
	// TODO: Preload参数"Book" "Information" 需要为结构体名称(大写就需要大写)
	err := global.MysqlDB.Preload("Book").Preload("Information").Where("id = ?", id).First(&student).Error

	return student, err
}

func (d *dao) FindAllStudent() ([]Student, error) {
	var students []Student
	//err := global.MysqlDB.Preload("book").Preload("Information").Find(&students).Error
	err := global.MysqlDB.Preload("Book").Find(&students).Error

	return students, err
}

func (d *dao) AddBook(book *Book) error {
	err := global.MysqlDB.Model(&Book{}).Create(book).Error
	return err
}

func (d *dao) RemoveBook(id uint) error {
	//global.MysqlDB.Model(&Book{}).Where("id = ?", id).Delete()
	err := global.MysqlDB.Model(&Book{}).Where("id = ?", id).Delete(&Book{}, id).Error

	return err
}

func (d *dao) AddLanguage(language *Language) error {
	err := global.MysqlDB.Model(&Language{}).Create(language).Error
	return err
}

func (d *dao) DeleteLanguage(id uint) error {
	err := global.MysqlDB.Model(&Language{}).Where("id = ?", id).Error

	return err
}

func (d *dao) BindStudentAndLanguage(studentId uint, languageIds []uint) error {
	global.MysqlDB.Transaction(func(tx *gorm.DB) error {

		//var languageList []Language
		//err := global.MysqlDB.Model(&Language{}).Where("id in (?)", languageIds).Find(&languageList).Error
		//if err != nil {
		//	return err
		//}
		var languageList []Language
		for _, v := range languageIds {
			languageList = append(languageList, Language{
				ID: v,
			})
		}

		var student Student
		err := global.MysqlDB.Model(&Student{}).Where("id = ? ", studentId).First(&student).Error
		if err != nil {
			return err
		}
		//
		//student.Language = languageList
		//err = global.MysqlDB.Model(&Student{}).Where("id = ?", studentId).Update("language", Language{}).Error
		//err = global.MysqlDB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&Student{}).Where("id = ?", student.ID).Save(&student).Error
		//err = global.MysqlDB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&Student{}).Where("id = ?", student.ID).Updates(&student).Error

		// INSERT INTO `language` (`name`,`id`) VALUES ('chinese',1),('english',3) ON DUPLICATE KEY UPDATE `id`=`id`
		// INSERT INTO `student_language` (`student_id`,`language_id`) VALUES (1,1),(1,3) ON DUPLICATE KEY UPDATE `student_id`=`student_id`
		// UPDATE `student` SET `name`='hanami' WHERE `id` = 1
		//err = global.MysqlDB.Updates(&student).Error

		//err = global.MysqlDB.Model(&Student{}).Where("student_id = ? ", studentId).Unscoped().Delete(&StudentLanguage{}).Error
		//if err != nil {
		//	return err
		//}

		//err = global.MysqlDB.Save(&student).Error
		// 替换关联
		//err = global.MysqlDB.Model(&Student{}).Where("id = ?", studentId).Association("Language").Replace(languageList)
		//err = global.MysqlDB.Model(&student).Association("Language").Replace(languageList)
		//err = global.MysqlDB.Model(&student).Omit("Language").Association("Language").Replace(languageList)
		err = global.MysqlDB.Model(&student).Association("Language").Replace(languageList)
		//global.MysqlDB.Model(&student).Where("id = ? ", studentId).Unscoped().Delete(&)
		return err
	})

	return nil
}
