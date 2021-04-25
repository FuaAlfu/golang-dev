package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	//model
	Product struct {
		gorm.Model
		Code  string
		Price uint
	}
)

var product Product

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "P1", Price: 100})

	db.First(&product, 1)
	db.First(&product, "code = ?", "P1")

	db.Model(&product).Update("Price", 8)
	db.Delete(&product, 1)

}
