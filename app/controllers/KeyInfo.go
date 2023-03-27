package controllers

import (
	"AuthService/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWK(c *gin.Context) {
	JwtInfo := new(service.JWTInfo)
	jwk := JwtInfo.InitJWK()
	c.JSON(http.StatusOK, jwk)
}
