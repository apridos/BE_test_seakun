package Routes

import (
	Controller "seakun/Controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// APIs
	grp_v1 := r.Group("/api/v1/")
	{
		grp_v1.POST("teacher/create", Controller.CreateTeacher)
		grp_v1.POST("teacher/edit", Controller.EditTeacher)
		grp_v1.POST("teacher/delete", Controller.DeleteTeacher)
		grp_v1.GET("teacher/all", Controller.GetAllTeacherData)
		grp_v1.GET("teacher/data/:id", Controller.GetTeacherDataById)
		grp_v1.GET("teacher/search/:name/:birth_date", Controller.SearchTeachers)

		grp_v1.POST("admin/login", Controller.AdminLogin)
	}

	return r
}
