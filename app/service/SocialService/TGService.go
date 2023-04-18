package SocialService

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

type TGService struct {
	conf *oauth2.Config
}

func NewTGService() *TGService {
	return &TGService{conf: &oauth2.Config{
		ClientID:     viper.GetString("VKONTAKTE_CLIENT_ID"),
		ClientSecret: viper.GetString("VKONTAKTE_CLIENT_SECRET"),
		Endpoint:     vk.Endpoint,
		RedirectURL:  viper.GetString("VKONTAKTE_REDIRECT_URI"),
		Scopes:       []string{"email", "phone_number"},
	}}
}
func (tgs TGService) index() string {
	return tgs.conf.AuthCodeURL("s")
}
