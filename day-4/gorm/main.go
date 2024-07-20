package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"unique"`
	Price uint
}

var db *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("failed to establish database connection:", err)
		return
	}
	log.Println("db connected")

}

func main() {
	//Migrate the schema
	err := db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal("Failed to migrate", err)
	}
	//createProduct("D2", 100)
	//
	//// Read
	//findProductByID(1)
	//
	//findAllProducts()
	//
	//// Update
	//updateProduct(5, "F422", 200)

	// Delete
	deleteProduct(1)
	fmt.Println("this is main function")
}

func createProduct(code string, price uint) {
	product := Product{Code: code, Price: price}

	err := db.Create(&product).Error
	if err != nil {
		log.Fatal("Failed to create product:", err)
	}
	log.Printf("Created product: %+v\n", product)
}

func findProductByID(id uint) {
	var product Product

	err := db.First(&product, id).Error
	if err != nil {
		log.Fatal("Failed to find product by ID:", err)
	}
	log.Printf("Found product by ID: %+v\n", product)
}

func findAllProducts() {
	var products []Product
	err := db.Find(&products).Error
	if err != nil {
		log.Fatal("Failed to find all products:", err)
	}
	log.Printf("Found all products: %+v\n", products)
}

func updateProduct(id uint, newCode string, newPrice uint) {
	var product Product
	err := db.First(&product, id).Error
	if err != nil {
		log.Fatal("Failed to find product for update:", err)
	}
	err = db.Model(&product).Updates(Product{Code: newCode, Price: newPrice}).Error
	if err != nil {
		log.Fatal("Failed to update product:", err)
	}
	log.Printf("Updated product: %+v\n", product)
}

func deleteProduct(id uint) {
	err := db.Delete(&Product{}, id).Error
	if err != nil {
		log.Fatal("Failed to delete product:", err)
	}
	log.Printf("Deleted product with ID: %d\n", id)
}
