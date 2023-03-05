package test_student

import (
	"github.com/gin-gonic/gin"
	"simple-video-server/core"
)

func SetupRoutes(v1 *gin.RouterGroup) {

	v1.POST("/student", core.ToHandler(Api.AddStudent, "student"))

	v1.GET("/student/:id", core.ToHandler(Api.GetInformation, "student"))

	v1.GET("/student", core.ToHandler(Api.GetAll, "student"))

	v1.POST("/student-information", core.ToHandler(Api.AddInformation, "student"))

	v1.POST("/student-book", core.ToHandler(Api.AddBook, "student"))

	v1.DELETE("/student-book/:id", core.ToHandler(Api.DeleteBook, "student"))

	v1.POST("/language", core.ToHandler(Api.AddLanguage, "language"))

	v1.DELETE("/language/:id", core.ToHandler(Api.DeleteLanguage, "language"))

	v1.POST("bind-student-language", core.ToHandler(Api.BindStudentAndLanguage, "student"))
}
