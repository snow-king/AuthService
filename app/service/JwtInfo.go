package service

import (
	"AuthService/app/errors"
	"AuthService/app/models"
	"encoding/json"
	"os"
)

type JWTInfo struct {
	JWK models.JWK
}

func (i JWTInfo) InitJWK() (models.JWK, error) {
	pwd, _ := os.Getwd()
	file, err := os.ReadFile(pwd + "/app/assets/symmetric.json")
	if err != nil {
		return i.JWK, errors.ErrFileDoesNotExist
	}
	err = json.Unmarshal(file, &i.JWK)
	return i.JWK, nil
}
