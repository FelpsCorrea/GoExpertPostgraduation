package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
	gorm.Model
}

// type SerialNumber struct {
// 	ID        int `gorm:"primaryKey"`
// 	Number    string
// 	ProductID int
// 	gorm.Model
// }

type Product struct {
	ID         int        `gorm:"primaryKey"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Categories []Category `gorm:"many2many:products_categories;"`
	// SerialNumber SerialNumber
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3307)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// Iniciar transação
	tx := db.Begin()
	var c Category

	// Está lockando a primeira categoria com limite 1 dizendo que será para um update
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error

	if err != nil {
		panic(err)
	}

	// Alterando a info
	c.Name = "Eletronicos"

	// Salvando a alteração para o commit
	tx.Debug().Save(&c)

	// Commit
	tx.Commit()

	// create category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// category2 := Category{Name: "Cozinha"}
	// db.Create(&category2)

	// // create product
	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.00,
	// 	Categories: []Category{category, category2},
	// })

	// // Inverso create category
	// products := []Product{
	// 	{Name: "Vara de Pesca", Price: 100.00},
	// 	{Name: "Barraca", Price: 100.00},
	// }

	// db.Create(&products)

	// db.Create(&Category{
	// 	Name:     "Lazer",
	// 	Products: products,
	// })

	// db.Create(&Product{
	// 	Name:       "Panela",
	// 	Price:      100.00,
	// 	CategoryID: 2,
	// })

	// var categories []Category

	// err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	fmt.Println(category.Name, ":")
	// 	for _, product := range category.Products {
	// 		println("- ", product.Name, category.Name)
	// 	}
	// }

	// db.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1,
	// })

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p.Name, p.Category.Name, p.SerialNumber.Number)
	// }

	// create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.00},
	// 	{Name: "Mouse", Price: 500.00},
	// 	{Name: "Keyboard", Price: 100.00},
	// }

	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 1)

	// db.First(&product, "name = ?", "Mouse")

	// select all
	// var products []Product

	// db.Find(&products)

	// db.Limit(2).Offset(2).Find(&products)

	// where
	// db.Where("price > ?", 100).Find(&products)

	// Like
	// db.Where("name LIKE ?", "%book%").Find(&products)

	// Salvar produto alterado
	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// Delete
	// var p2 Product
	// db.First(&p2, 1)
	// db.Delete(&p2)
}
