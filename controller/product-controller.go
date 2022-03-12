package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/doffy007/go-api-jwt/dto"
	"github.com/doffy007/go-api-jwt/entity"
	"github.com/doffy007/go-api-jwt/helper"
	"github.com/doffy007/go-api-jwt/service"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	FindByTitle(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type productController struct {
	productService service.ProductService
	jwtService     service.JWTService
}

//create new instacnce of product controller
func NewProductController(prodServ service.ProductService, jwtServ service.JWTService) ProductController {
	return &productController{
		productService: prodServ,
		jwtService:     jwtServ,
	}
}

func (c *productController) All(context *gin.Context) {
	var products []entity.Product = c.productService.All()
	res := helper.BuildResponse(true, "OK!", nil, products)
	context.JSON(http.StatusOK, res)
}

func (c *productController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("ID not found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var product entity.Product = c.productService.FindByID(id)
	if (product == entity.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "Nothing data can find", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", nil, product)
		context.JSON(http.StatusOK, res)
	}
}

func (c *productController) FindByTitle(context *gin.Context) {
	var title entity.Product = c.productService.FindByTitle(context.Param("title"))
	res := helper.BuildResponse(true, "OK!", nil, title)
	context.JSON(http.StatusOK, res)
}

func (c *productController) Insert(context *gin.Context) {
	var productCreateDTO dto.ProductCreateDTO
	errDTO := context.ShouldBind(&productCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Cannot Build Product", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			productCreateDTO.UserID = convertUserID
		}
		res := c.productService.Insert(productCreateDTO)
		response := helper.BuildResponse(true, "OK!", nil, res)
		context.JSON(http.StatusOK, response)
	}
}

func (c *productController) Update(context *gin.Context) {
	var productUpdateDTO dto.ProductUpdateDTO
	errDTO := context.ShouldBind(&productUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.productService.IsAllowedToEdit(userID, uint(productUpdateDTO.ID)) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			productUpdateDTO.UserID = id
		}
		result := c.productService.Update(productUpdateDTO)
		response := helper.BuildResponse(true, "OK", nil, result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "youre not owner account", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *productController) Delete(context *gin.Context) {
	var product entity.Product
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed cannot get id", "No param id found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	product.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.productService.IsAllowedToEdit(userID, uint(product.ID)) {
		c.productService.Delete(product)
		res := helper.BuildResponse(true, "Delete", nil, helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont has permission", "youre not owner this accout", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *productController) getUserIDByToken(token string) string {
	t, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := t.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
