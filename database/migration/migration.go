package migration

import (
	"fiber-go/model/entity"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func RunMigrate() {
// 	database.DB.AutoMigrate(&entity.User{})
// }

func Migration() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go-fiber?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Something wrong", err)
		return nil
	}

	fmt.Println("Database connected")
	
	db.AutoMigrate(&entity.User{})
	db.Debug()

	data_1 := []entity.User{
		{Name: "Afif", Email: "afif@gmail.com", Phone: "123456789", Address: "Address 1"},
		{Name: "Joko", Email: "joko@gmail.com", Phone: "987654321", Address: "Address 2"},
		{Name: "Eko", Email: "eko@gmail.com", Phone: "555555555", Address: "Address 3"},
		{Name: "Bambang", Email: "bambang@gmail.com", Phone: "444444444", Address: "Address 4"},
	}

	// Create records
	if err := db.Create(&data_1).Error; err != nil {
		fmt.Println("Error inserting data:", err)
	}

	return db
}