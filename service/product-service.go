package service

import (
	"fmt"
	"log"

	"github.com/doffy007/go-api-jwt/dto"
	"github.com/doffy007/go-api-jwt/entity"
	"github.com/doffy007/go-api-jwt/repository"
	"github.com/mashingan/smapping"
)

//ProductService is contract about something ProductService can do
type ProductService interface {
	Insert(product dto.ProductCreateDTO) entity.Product
	Update(product dto.ProductUpdateDTO) entity.Product
	Delete(product entity.Product)
	All() []entity.Product
	FindByID(productID uint64) entity.Product
	FindByTitle(productTitle string) entity.Product
	IsAllowedToEdit(userID string, productID uint) bool
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (service *productService) Insert(p dto.ProductCreateDTO) entity.Product {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.InsertProduct(product)
	return res
}

func (service *productService) Update(p dto.ProductUpdateDTO) entity.Product {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.productRepository.UpdateProduct(product)
	return res
}


func (service *productService) Delete(product entity.Product) {
	service.productRepository.DeleteProduct(product)
}

func (service *productService) All() []entity.Product {
	return service.productRepository.AllProduct()
}

func (service *productService) FindByID(bookID uint64) entity.Product {
	return service.productRepository.FindProductByID(bookID)
}

func (sevice *productService) FindByTitle(productTitle string) entity.Product {
	return sevice.productRepository.FindProductByTitle(productTitle)
}

func (service *productService) IsAllowedToEdit(userID string, productID uint) bool {
	p := service.productRepository.FindProductByID(uint64(productID))
	id := fmt.Sprintf("%v", p.UserID)
	return userID == id
}