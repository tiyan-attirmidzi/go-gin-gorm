package repositories

import (
	"github.com/tiyan-attirmidzi/go-rest-api/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	StoreUser(user entities.User) entities.User
	UpdateUser(user entities.User) entities.User
	VerifyCredential(email, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entities.User
	ProfileUser(userID string) entities.User
}

type userConnectin struct {
	connection *gorm.DB
}

func NewUserRespository(db *gorm.DB) UserRepository {
	return &userConnectin{
		connection: db,
	}
}

func (db *userConnectin) StoreUser(user entities.User) entities.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnectin) UpdateUser(user entities.User) entities.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnectin) VerifyCredential(email, password string) interface{} {
	var user entities.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnectin) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entities.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnectin) FindByEmail(email string) entities.User {
	var user entities.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnectin) ProfileUser(userID string) entities.User {
	var user entities.User
	db.connection.Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		// log.Fatal("Generate Hash Failed!")
		panic("Failed to Hash Password!")
	}
	return string(hash)
}
