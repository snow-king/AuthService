package service

import (
	"AuthService/app/models"
	"encoding/base64"
	"github.com/spf13/viper"
)

type JWTInfo struct {
	JWK models.JWK
}

func (i JWTInfo) InitJWK() models.JWK {
	key := viper.GetString("SIGNING_KEY")
	i.JWK.Keys = append(i.JWK.Keys, struct {
		Kty string `json:"kty,omitempty"`
		Kid string `json:"kid,omitempty"`
		K   string `json:"k,omitempty"`
		Alg string `json:"alg,omitempty"`
		Use string `json:"use,omitempty"`
	}{Kty: "oct", Kid: viper.GetString("KID"), K: key, Alg: "HS256"})
	return i.JWK
}
