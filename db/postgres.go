package db

import (
	"fmt"
	"github.com/Mirzoev-Parviz/movie-recommendation/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func initDB() *gorm.DB {
	settingsParam := utils.AppSettings.DBSettings

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Dushanbe",
		settingsParam.Host, settingsParam.User, settingsParam.Password, settingsParam.Database,
		settingsParam.Port, settingsParam.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[db] failed to connect to postgreSQL database")
	}

	return db
}

func StartDbConnection() {
	log.Println("[db] connected to database")
	database = initDB()
}

func GetDBConn() *gorm.DB {
	return database
}

func DisconnectDB(db *gorm.DB) {
	_db, err := db.DB()
	if err != nil {
		log.Fatal("[db] failed to kill connection from database. Error is: ", err.Error())
	}

	_db.Close()
	log.Println("[db] successfully disconnected from database")
}
