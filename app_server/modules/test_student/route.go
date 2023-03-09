package test_student

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	// add
	v1.POST("/student", core.ToHandler(Api.AddStudent, "student"))
	// get by id
	v1.GET("/student/:id", core.ToHandler(Api.Get, "student"))
	// get list
	v1.GET("/student", core.ToHandler(Api.GetAll, "student"))
	// add information
	v1.POST("/student-information", core.ToHandler(Api.AddInformation, "student"))

	// add book
	v1.POST("/student-book", core.ToHandler(Api.AddBook, "student"))
	// get books
	v1.GET("/student-book", core.ToHandler(Api.GetBooks, "student"))
	// delete book
	v1.DELETE("/student-book/:id", core.ToHandler(Api.DeleteBook, "student"))

	// add language
	v1.POST("/language", core.ToHandler(Api.AddLanguage, "language"))
	// delete language
	v1.DELETE("/language/:id", core.ToHandler(Api.DeleteLanguage, "language"))
	// bind language to student
	v1.POST("/bind-student-language", core.ToHandler(Api.BindStudentAndLanguage, "student"))
}
