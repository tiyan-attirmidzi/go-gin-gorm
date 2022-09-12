package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/tiyan-attirmidzi/go-rest-api/dto"
	"github.com/tiyan-attirmidzi/go-rest-api/entities"
	"github.com/tiyan-attirmidzi/go-rest-api/repositories"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entities.User
	Profile(userID string) entities.User
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Update(user dto.UserUpdateDTO) entities.User {
	userToUpdate := entities.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updateUser := s.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (s *userService) Profile(userID string) entities.User {
	return s.userRepository.ProfileUser(userID)
}
