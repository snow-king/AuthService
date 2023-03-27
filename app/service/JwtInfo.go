package service

import (
	"AuthService/app/models"
)

type JWTInfo struct {
	JWK models.JWK
}

func (i JWTInfo) InitJWK() models.JWK {
	i.JWK.Keys = append(i.JWK.Keys, struct {
		Kty string `json:"kty,omitempty"`
		Kid string `json:"kid,omitempty"`
		Alg string `json:"alg,omitempty"`
	}{Kty: "HS", Kid: "NjVBRjY5MDlCMUIwNzU4RTA2QzZFMDQ4QzQ2MDAyQjVDNjk1RTM2Qg", Alg: "HS256"})
	return i.JWK
}
