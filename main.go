//main.go
package main

import (
	"fmt"
	Config "seakun/Config"
	Model "seakun/Model"
	"seakun/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Model.Admin{})

	r := Routes.SetupRouter()
	//running
	r.Run(":8085")
}
