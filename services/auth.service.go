package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/tiyan-attirmidzi/go-gin-gorm/dto"
	"github.com/tiyan-attirmidzi/go-gin-gorm/entities"
	"github.com/tiyan-attirmidzi/go-gin-gorm/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(user dto.AuthSignUpDTO) entities.User
	FindByEmail(email string) entities.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repositories.UserRepository
}

type Err struct {
	Field string
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (s *authService) VerifyCredential(email, password string) interface{} {
	res := s.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entities.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return Err{Field: "password"}
	}
	return Err{Field: "email"}
	// return false
}

func (s *authService) CreateUser(user dto.AuthSignUpDTO) entities.User {
	userToCreate := entities.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := s.userRepository.StoreUser(userToCreate)
	return res
}

func (s *authService) FindByEmail(email string) entities.User {
	return s.userRepository.FindByEmail(email)
}

func (s *authService) IsDuplicateEmail(email string) bool {
	res := s.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		// log.Fatal(err)
		return false
	}
	return true
}
