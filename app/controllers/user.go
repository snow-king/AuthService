package controllers

import (
	"AuthService/app/errors"
	"AuthService/app/service"
	"AuthService/app/structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AppendRole(c *gin.Context) {
	var inp structures.UserRole
	if err := c.Bind(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userS, err := service.NewUserService(inp.UserId)
	roles := userS.AppendRole(inp.RoleId)
	if err != nil {
		if err == errors.ErrUserDoesNotExist {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, err.Error(), ""))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, newSignInResponse(StatusError, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, roles)
}
func AppendSocNetworks(c *gin.Context) {
	var inp structures.UserSocNetwork
	if err := c.Bind(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userS, err := service.NewUserService(inp.UserId)
	networks, err := userS.AddSocNetwork(inp.NetworkId, inp.NetworkNameId)
	if err != nil {
		if err == errors.ErrUserDoesNotExist {
			c.AbortWithStatusJSON(http.StatusBadRequest, newSignInResponse(StatusError, err.Error(), ""))
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, newSignInResponse(StatusError, err.Error(), ""))
		return
	}
	c.JSON(http.StatusCreated, networks)
}
