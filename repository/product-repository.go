package repository

import (
	"github.com/doffy007/go-api-jwt/entity"
	"gorm.io/gorm"
)

//ProductRepository is contract what ProductRepository can do to db
type ProductRepository interface {
	InsertProduct(p entity.Product) entity.Product
	UpdateProduct(p entity.Product) entity.Product
	DeleteProduct(p entity.Product) 
	AllProduct() []entity.Product  //public can acess
	FindProductByID(productID uint64) entity.Product
	FindProductByTitle(productTitle string) entity.Product
}

type productConnection struct {
	connection *gorm.DB
}
//NewProductRepository create new instance of productRepository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) InsertProduct(p entity.Product) entity.Product {
	db.connection.Save(&p)
	db.connection.Preload("User").Find(&p)
	return p
}

func (db *productConnection) UpdateProduct(p entity.Product) entity.Product {
	db.connection.Save(&p)
	db.connection.Preload("User").Find(&p)
	return p
}

func (db *productConnection) DeleteProduct(b entity.Product) {
	db.connection.Delete(&b)
}

func (db *productConnection) FindProductByID(productID uint64) entity.Product {
	var product entity.Product
	db.connection.Preload("User").Find(&product, productID)
	return product
}

func (db *productConnection) FindProductByTitle(productTitle string) entity.Product {
	var product entity.Product
	db.connection.Preload("User").Find(&product, productTitle)
	return product
}

func (db *productConnection) AllProduct() []entity.Product {
	var products []entity.Product
	db.connection.Preload("User").Find(&products)
	return products
}