package database

import (
	"fmt"

		"gorm.io/driver/mysql"
		"gorm.io/gorm"	
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/go-sql?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("cannot connect to database", err)
	}

	fmt.Println("connected to database")
}
