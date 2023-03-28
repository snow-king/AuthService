package service

import (
	"AuthService/app/auth"
	"AuthService/app/errors"
	"AuthService/app/models"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
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

// SignUp - регистрация пользователя в системе
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
	if err := models.DbEIS.Where(&models.User{
		NdsLogin: userAuth.Login,
		Password: password,
	}).First(&user).Error; err != nil {
		log.Errorf("error on inserting user: %s", err.Error())
		return "", errors.ErrUserDoesNotExist
	}
	var roles []string
	models.DbEIS.Model(&models.User{}).Preload("Roles").Find(&user)
	for _, role := range user.Roles {
		fmt.Println(role.JwtName)
		roles = append(roles, role.JwtName)
	}
	fmt.Println(roles)
	if !slices.Contains(roles, userAuth.Role) {
		return "", errors.ErrUserDoesNotHaveAccess
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
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
