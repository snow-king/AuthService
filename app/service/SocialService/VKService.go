package SocialService

import (
	"AuthService/app/models"
	"AuthService/app/service"
	"fmt"
	vkApi "github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"log"
	"time"
)

type VkService struct {
	conf *oauth2.Config
}

func NewVkService() *VkService {
	return &VkService{conf: &oauth2.Config{
		ClientID:     viper.GetString("VKONTAKTE_CLIENT_ID"),
		ClientSecret: viper.GetString("VKONTAKTE_CLIENT_SECRET"),
		Endpoint:     vk.Endpoint,
		RedirectURL:  viper.GetString("VKONTAKTE_REDIRECT_URI"),
		Scopes:       []string{"email", "phone_number"},
	}}
}

func (vkS VkService) Index() string {
	return vkS.conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (vkS VkService) Callback(code string) (string, error) {
	ctx := context.Background()
	tok, err := vkS.conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	email := fmt.Sprintf("%v", tok.Extra("email"))
	id := fmt.Sprintf("%v", tok.Extra("user_id"))
	// создаем клиент для получения данных из API VK
	//client := vkApi.NewVK(tok.AccessToken)
	login := models.LoginUser{
		Login:    email,
		Password: "",
		Role:     "",
	}
	var AuthUseCase = service.NewAuthorizer(
		viper.GetString("HASH_SALT"),
		[]byte(viper.GetString("SIGNING_KEY")),
		viper.GetDuration("TOKEN_TTL")*time.Second,
	)
	user, token, err := AuthUseCase.SignUp(login)
	if err != nil {
		return "", err
	}
	userService, err := service.NewUserService(user.Id)
	if err != nil {
		return "", err
	}
	_, err = userService.AddSocNetwork(id, 1)
	if err != nil {
		return "", err
	}
	return token, err
}

// получение доп информации о пользователе
func getCurrentUser(api *vkApi.VK) object.UsersUser {
	users, err := api.UsersGet(vkApi.Params{
		"fields": "sex, city, country",
	})
	if err != nil {
		log.Fatal(err)
	}
	return users[0]
}
