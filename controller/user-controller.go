package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/doffy007/go-api-jwt/dto"
	"github.com/doffy007/go-api-jwt/helper"
	"github.com/doffy007/go-api-jwt/service"
	"github.com/gin-gonic/gin"
)

//UserController is a contract interface can do
type UserController interface {
	Update(contex *gin.Context)
	Profile(contex *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

//NewUserController creating an new instance of UserController
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", nil, u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK!", nil, user)
	context.JSON(http.StatusOK, res)
}
