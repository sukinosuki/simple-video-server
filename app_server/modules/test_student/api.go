package test_student

import (
	"simple-video-server/core"
	"strconv"
)

type _api struct {
}

var Api = &_api{}

func (api *_api) AddStudent(c *core.Context) (uint, error) {

	var studentAdd StudentAdd
	err := c.ShouldBind(&studentAdd)
	if err != nil {
		panic(err)
	}

	student := Student{
		Name: studentAdd.Name,
	}
	err = Dao.AddStudent(&student)
	if err != nil {
		panic(err)
	}

	return student.ID, nil
}

func (api *_api) AddInformation(c *core.Context) (uint, error) {
	var informationAdd InformationAdd
	err := c.ShouldBind(&informationAdd)
	if err != nil {
		panic(err)
	}
	information := Information{
		Age:       informationAdd.Age,
		StudentID: informationAdd.StudentID,
	}
	err = Dao.AddInformation(&information)
	if err != nil {
		panic(err)
	}

	return information.ID, err
}

func (api *_api) Get(c *core.Context) (any, error) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		panic(err)
	}

	student, err := Dao.FindStudentById2(uint(id))

	if err != nil {
		panic(err)
	}

	return student, err
}

func (api *_api) GetAll(c *core.Context) ([]Student, error) {
	students, err := Dao.FindAllStudent()
	if err != nil {
		panic(err)
	}

	return students, err
}

func (api *_api) GetBooks(c *core.Context) ([]BookSimple, error) {
	books, err := Dao.GetStudentBooks(4)
	if err != nil {
		panic(err)
	}

	return books, err
}

func (api *_api) AddBook(c *core.Context) (bool, error) {
	var add BookAdd
	err := c.ShouldBind(&add)
	if err != nil {
		panic(err)
	}
	book := Book{
		Name:      add.Name,
		StudentID: add.StudentID,
	}
	err = Dao.AddBook(&book)
	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *_api) DeleteBook(c *core.Context) (bool, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	err = Dao.RemoveBook(uint(id))
	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *_api) AddLanguage(c *core.Context) (uint, error) {
	var add LanguageAdd
	err := c.ShouldBind(&add)
	if err != nil {
		panic(err)
	}
	language := Language{
		Name: add.Name,
	}
	err = Dao.AddLanguage(&language)
	if err != nil {
		panic(err)
	}

	return language.ID, err
}

func (api *_api) DeleteLanguage(c *core.Context) (bool, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	err = Dao.DeleteLanguage(uint(id))

	if err != nil {
		panic(err)
	}

	return true, err
}

func (api *_api) BindStudentAndLanguage(c *core.Context) (bool, error) {

	var data BindStudentAndLanguage
	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err)
	}

	err = Dao.BindStudentAndLanguage2(data.StudentID, data.LanguageID)
	if err != nil {
		panic(err)
	}

	return true, err
}
