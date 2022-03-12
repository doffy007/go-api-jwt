package service

import (
	"log"

	"github.com/doffy007/go-api-jwt/dto"
	"github.com/doffy007/go-api-jwt/entity"
	"github.com/doffy007/go-api-jwt/repository"
	"github.com/mashingan/smapping"
)

//UserService is contract about something UserService can do
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

//NewUserService create a new instance of UserService
func NewUserservice(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v :", err)
	}
	updateUser := service.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
