package controllers

import (
	"AuthService/app/errors"
	"AuthService/app/service/SocialService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func choseSocial(name string) SocialService.SocService {
	var network SocialService.SocService
	switch name {
	case "vk":
		network = SocialService.NewVkService()
	default:
		return nil
	}
	return network
}
func Index(c *gin.Context) {
	socNetworkName := c.Param("service")
	network := choseSocial(socNetworkName)
	if network != nil {
		url := network.Index()
		c.Redirect(http.StatusMovedPermanently, url)
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, "не верный сервис авторизации", ""))
	}
}
func Callback(c *gin.Context) {
	socNetworkName := c.Param("service")
	network := choseSocial(socNetworkName)
	authCode := c.Request.URL.Query()["code"]
	token, err := network.Callback(authCode)
	if err != nil {
		if err == errors.ErrInvalidAccessToken {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, err.Error(), ""))
			return
		}
		if err == errors.ErrUserDoesNotHaveAccess {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, err.Error(), ""))
			return
		}
		if err == errors.ErrUserDoesNotExist {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, err.Error(), ""))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, newSignInResponse(StatusError, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, newSignInResponse(StatusOk, "", token))
}
