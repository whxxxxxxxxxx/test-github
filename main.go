package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func create(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

func insert(db *gorm.DB) {
	p := Product{Code: "D42", Price: 100}
	db.Create(&p)
}

func update(db *gorm.DB) {
	var product Product
	db.First(&product, 1)
	db.Model(&product).Update("Price", 1000)
}

func find(db *gorm.DB) {
	var product Product
	db.First(&product, 1) // find product with integer primary key
	fmt.Printf("product: %v\n", product)
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	update(db)

}
