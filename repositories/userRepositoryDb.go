package repositories

import (
	"app/dto"
	"app/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

type UserRepository interface {
	Create(userDto dto.UserRequestDto) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

func (h UserRepositoryDb) Create(userDto dto.UserRequestDto) (models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.MinCost)

	if err != nil {
		fmt.Println("HASH ERROR")
		os.Exit(1)
	}

	user := models.User{
		Email:    userDto.Email,
		Name:     userDto.Name,
		Password: string(hash),
	}
	fmt.Println(user)

	result := h.Db.Create(&user)

	return user, result.Error
}

func (h UserRepositoryDb) CheckUser(email string, password string) (models.User, error) {
	user, err := h.FindByEmail(email)

	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (h UserRepositoryDb) FindByEmail(email string) (models.User, error) {
	var user = models.User{}

	result := h.Db.Where("email = ?", email).First(&user)

	return user, result.Error
}

func NewUserRepositoryDb(db *gorm.DB) UserRepositoryDb {
	return UserRepositoryDb{Db: db}
}
