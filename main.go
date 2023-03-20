package main

import (
	server "AuthService/app"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	app := new(server.App)
	if err := app.Run("8080"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
