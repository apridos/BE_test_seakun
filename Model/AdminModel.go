package model

import (
	Config "seakun/Config"

	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id       int    `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `gorm:"size:25"`
}

func (b *Admin) TableName() string {
	return "admin"
}

func FindAdminByUsername(admin *Admin, username string) (err error) {
	if err = Config.DB.Where("username = ?", username).First(admin).Error; err != nil {
		return err
	}
	return nil
}
