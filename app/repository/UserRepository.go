package repository

import (
	"AuthService/app/errors"
	"AuthService/app/models"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=IUserRepository
type IUserRepository interface {
	Create(userAuth models.LoginUser) (models.User, error)
	Get(userAuth models.LoginUser) (models.User, error)
	GetRoles(user models.User) []string
}

type Repository struct {
	Database *gorm.DB
}

var validate *validator.Validate

// Create создание записи user
func (u Repository) Create(userAuth models.LoginUser) (models.User, error) {
	user := models.User{
		NdsLogin: userAuth.Login,
		Password: userAuth.Password,
	}
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return models.User{}, errors.ErrFiledCheckData
	}
	if err := u.Database.Where(user).FirstOrCreate(&user).Error; err != nil {
		log.Errorf("error on inserting user: %s", err.Error())
		return models.User{}, errors.ErrUserDoesNotExist
	}
	return user, nil
}

// Get получение записи из бд
func (u Repository) Get(userAuth models.LoginUser) (models.User, error) {
	var user models.User
	validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		return models.User{}, errors.ErrFiledCheckData
	}
	if err := u.Database.Where(&models.User{
		NdsLogin: userAuth.Login,
		Password: userAuth.Password,
	}).First(&user).Error; err != nil {
		log.Errorf("error on inserting user: %s", err.Error())
		return models.User{}, errors.ErrUserDoesNotExist
	}
	return user, nil
}

// GetRoles получение ролей юзера
func (u Repository) GetRoles(user models.User) []string {
	var roles []string
	u.Database.Model(&models.User{}).Preload("Roles").Find(&user)
	for _, role := range user.Roles {
		roles = append(roles, role.JwtName)
	}
	return roles
}

func (u Repository) Delete() {

}
