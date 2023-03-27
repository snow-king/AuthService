package service

import (
	"AuthService/app/models"
	"github.com/spf13/viper"
)

type JWTInfo struct {
	JWK models.JWK
}

func (i JWTInfo) InitJWK() models.JWK {
	i.JWK.Keys = append(i.JWK.Keys, struct {
		Kty string `json:"kty,omitempty"`
		Kid string `json:"kid,omitempty"`
		K   string `json:"k,omitempty"`
		Alg string `json:"alg,omitempty"`
	}{Kty: "oct", Kid: viper.GetString("kid"), K: viper.GetString("signing_key"), Alg: "HS256"})
	return i.JWK
}
