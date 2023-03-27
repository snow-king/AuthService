package controllers

import (
	"AuthService/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWK(c *gin.Context) {
	JwtInfo := new(service.JWTInfo)
	jwk, err := JwtInfo.InitJWK()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, newSignInResponse(StatusError, err.Error(), ""))
		return
	}
	c.JSON(http.StatusOK, jwk)
}
