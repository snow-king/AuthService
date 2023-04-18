package service

import (
	"AuthService/app/models"
	"AuthService/app/repository"
	"AuthService/app/structures"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
	"time"
)

type Authorizer struct {
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
	userRepo       repository.Repository
}

func NewAuthorizer(hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
		userRepo:       repository.Repository{Database: models.DbEIS},
	}
}

// SignUp - регистрация пользователя в системе
func (a *Authorizer) SignUp(userAuth models.LoginUser) (models.User, string, error) {
	// Create password hash
	pwd := md5.New()
	pwd.Write([]byte(userAuth.Password))
	userAuth.Password = fmt.Sprintf("%x", pwd.Sum(nil))
	user, err := a.userRepo.Create(userAuth)
	if err != nil {
		return user, "", err
	}
	token, err := a.generateToken(user)
	return user, token, err
}

// SignIn авторизация пользователя по логину и паролю
func (a *Authorizer) SignIn(userAuth models.LoginUser) (string, error) {
	pwd := md5.New()
	pwd.Write([]byte(userAuth.Password))
	//pwd.Write([]byte(a.hashSalt))
	userAuth.Password = fmt.Sprintf("%x", pwd.Sum(nil))
	user, err := a.userRepo.Create(userAuth)
	if err != nil {
		return "", err
	}
	return a.generateToken(user)
}

// generateToken Генерация токена по модели юзера
func (a *Authorizer) generateToken(user models.User) (string, error) {
	roles := a.userRepo.GetRoles(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &structures.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: user.NdsLogin,
		UserId:   user.Id,
		Roles:    roles,
	})
	token.Header["kid"] = viper.GetString("kid")
	return token.SignedString(a.signingKey)
}
