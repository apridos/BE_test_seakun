package model

import (
	Config "seakun/Config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Teacher struct {
	Id        int        `json:"id"`
	Name      string     `json:"name" gorm:"size:30"`
	BirthDate *time.Time `json:"birth_date" gorm:"not null"`
}

func (b *Teacher) TableName() string {
	return "teacher"
}

func CreateTeacher(teacher *Teacher) (err error) {
	if err = Config.DB.Create(teacher).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTeacher(teacher *Teacher) (err error) {
	Config.DB.Save(teacher)
	return nil
}

func FindTeacherById(teacher *Teacher, teacherId string) (err error) {
	if err = Config.DB.Where("id = ?", teacherId).First(teacher).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTeacher(teacher *Teacher, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(teacher)
	return nil
}

func SearchTeacher(teachers *[]Teacher, name string, birthDate string) (err error) {
	if err = Config.DB.Where("name like ? and DATE(birth_date) = DATE(?)", "%"+name+"%", birthDate).Find(teachers).Error; err != nil {
		return err
	}
	return nil
}

func AllTeacher(teachers *[]Teacher) (err error) {
	if err = Config.DB.Find(teachers).Error; err != nil {
		return err
	}
	return nil
}
