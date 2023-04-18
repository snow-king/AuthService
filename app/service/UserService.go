package service

import (
	"AuthService/app/errors"
	"AuthService/app/models"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	User models.User
}

//go:generate go run github.com/vektra/mockery/v2@latest --name=IUserService
type IUserService interface {
	AppendRole(roleId int) []models.Roles
	DeleteRole(roleID int)
	AddSocNetwork(networkId string, networkNameID int) ([]models.UserNetworks, error)
}

func NewUserService(id int) (*UserService, error) {
	var user models.User
	if err := models.DbEIS.Where(&models.User{
		Id: id,
	}).Preload("Roles").Preload("UserNetworks").First(&user).Error; err != nil {
		log.Errorf("error on inserting user: %s", err.Error())
		return nil, errors.ErrUserDoesNotExist
	}
	return &UserService{User: user}, nil
}

func (s UserService) AppendRole(roleId int) []models.Roles {
	record := models.UserRoles{
		UserId: s.User.Id,
		RolId:  roleId,
	}
	models.DbEIS.Where(record).FirstOrCreate(&record)
	return s.User.Roles
}
func (s UserService) DeleteRole(roleID int) error {
	err := models.DbEIS.Where(&models.UserRoles{
		UserId: s.User.Id,
		RolId:  roleID,
	}).Delete(models.UserRoles{}).Error
	if err != nil {
		return err
	}
	return nil
}

// AddSocNetwork  Добавление соц сети к аккаунту
func (s UserService) AddSocNetwork(networkId string, networkNameID int) ([]models.UserNetworks, error) {
	record := models.UserNetworks{
		UserID:        s.User.Id,
		NetworkNameId: networkNameID,
		NetworkId:     networkId,
	}
	models.DbEIS.Where(record).FirstOrCreate(&record)
	return s.User.UserNetworks, nil
}
