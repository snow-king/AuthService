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
	key := base64.StdEncoding.EncodeToString([]byte(viper.GetString("signing_key")))
	i.JWK.Keys = append(i.JWK.Keys, struct {
		Kty string `json:"kty,omitempty"`
		Kid string `json:"kid,omitempty"`
		K   string `json:"k,omitempty"`
		Alg string `json:"alg,omitempty"`
	}{Kty: "oct", Kid: viper.GetString("kid"), K: key, Alg: "HS256"})
	return i.JWK
}
