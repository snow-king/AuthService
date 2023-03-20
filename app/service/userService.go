package service

import (
	"AuthService/app/auth"
	"AuthService/app/errors"
	"AuthService/app/models"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	log "github.com/sirupsen/logrus"
	"time"
)

type Authorizer struct {
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthorizer(hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}
func (a *Authorizer) SignUp(userAuth models.LoginUser) {
	// Create password hash
	pwd := md5.New()
	pwd.Write([]byte(userAuth.Password))
	pwd.Write([]byte(a.hashSalt))
	userAuth.Password = fmt.Sprintf("%x", pwd.Sum(nil))
}

// SignIn авторизация пользователя по логину и паролю
func (a *Authorizer) SignIn(userAuth models.LoginUser) (string, error) {
	pwd := md5.New()
	pwd.Write([]byte(userAuth.Password))
	//pwd.Write([]byte(a.hashSalt))
	password := fmt.Sprintf("%x", pwd.Sum(nil))
	var user models.User
	fmt.Println(password)
	if err := models.DbEIS.Where(&models.User{
		NdsLogin: userAuth.Login,
		Password: password,
	}).First(&user).Error; err != nil {
		log.Errorf("error on inserting user: %s", err.Error())
		return "", errors.ErrUserDoesNotExist
	}
	fmt.Println(jwt.At(time.Now().Add(a.expireDuration)))
	fmt.Println(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: user.NdsLogin,
		UserId:   user.Id,
	})
	return token.SignedString(a.signingKey)
}
