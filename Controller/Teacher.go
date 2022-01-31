package controller

import (
	"fmt"
	"net/http"
	"reflect"
	model "seakun/Model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {
	if !Authorize(c) {
		return
	}

	var teacher model.Teacher
	c.BindJSON(&teacher)

	fmt.Println(reflect.TypeOf(teacher.BirthDate))

	err := model.CreateTeacher(&teacher)
	if err != nil {
		c.Data(http.StatusBadRequest, "application/json", []byte(`{"result":"bad request"}`))
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result":"success"}`))
}

func EditTeacher(c *gin.Context) {
	if !Authorize(c) {
		return
	}

	var teacher model.Teacher
	c.BindJSON(&teacher)
	fmt.Println(teacher)

	err := model.UpdateTeacher(&teacher)

	if err != nil {
		fmt.Println(err)

		c.Data(http.StatusBadRequest, "application/json", []byte(`{"result":"bad request"}`))
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result":"success"}`))
}

func GetTeacherDataById(c *gin.Context) {
	if !Authorize(c) {
		return
	}

	var teacher model.Teacher
	teacherId := c.Params.ByName("id")

	err := model.FindTeacherById(&teacher, teacherId)

	if err != nil {
		c.Data(http.StatusNotFound, "application/json", []byte(`{"result":"not found"}`))
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func DeleteTeacher(c *gin.Context) {
	if !Authorize(c) {
		return
	}
	var teacher model.Teacher
	c.BindJSON(&teacher)

	err := model.DeleteTeacher(&teacher, strconv.Itoa(teacher.Id))
	if err != nil {
		c.Data(http.StatusBadRequest, "application/json", []byte(`{"result":"bad request"}`))
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(`{"result":"success"}`))
}

func SearchTeachers(c *gin.Context) {
	if !Authorize(c) {
		return
	}

	var teachers []model.Teacher
	err := model.SearchTeacher(&teachers, c.Params.ByName("name"), c.Params.ByName("birth_date"))
	if err != nil {
		c.Data(http.StatusBadRequest, "application/json", []byte(`{"result":"bad request"}`))
		return
	}

	c.JSON(http.StatusOK, teachers)
}

func GetAllTeacherData(c *gin.Context) {
	if !Authorize(c) {
		return
	}
	var teachers []model.Teacher

	err := model.AllTeacher(&teachers)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, teachers)
}
