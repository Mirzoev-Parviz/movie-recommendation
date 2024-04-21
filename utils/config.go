package utils

import (
	"github.com/Mirzoev-Parviz/movie-recommendation/models"
	"github.com/spf13/viper"
	"log"
)

var AppSettings models.Settings

func ReadSettings() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read `config.yaml` file. Error is: ", err.Error())
	}
	setup()
	dbSetup()
}

func setup() {
	AppSettings.AppParams.PortRun = viper.GetString("port")
}

func dbSetup() {
	AppSettings.DBSettings.User = viper.GetString("db.username")
	AppSettings.DBSettings.Database = viper.GetString("db.name")
	AppSettings.DBSettings.Password = viper.GetString("db.password")
	AppSettings.DBSettings.Host = viper.GetString("db.host")
	AppSettings.DBSettings.Port = viper.GetString("db.port")
	AppSettings.DBSettings.SSLMode = viper.GetString("db.sslmode")
}
