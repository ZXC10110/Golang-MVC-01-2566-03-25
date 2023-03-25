package main

import (
	"fmt"

	"mvc_63050096_2565_2/Config"
	Routes "mvc_63050096_2565_2/Controllers"
	"mvc_63050096_2565_2/Models"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	fmt.Println("Run Swagger : http://localhost:8080/swagger/index.html")
	fmt.Println()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.ChatCSIfElse{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
