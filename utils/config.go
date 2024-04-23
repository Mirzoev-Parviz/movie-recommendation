package utils

import (
	"github.com/Mirzoev-Parviz/movie-recommendation/internal/dto"
	"github.com/Mirzoev-Parviz/movie-recommendation/models"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

var AppSettings models.Settings

func ReadSettings() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read `config.yaml` file. Error is: ", err.Error())
	}
	setup()
	loadData()
	//dbSetup()
}

func setup() {
	AppSettings.AppParams.PortRun = viper.GetString("port")
}

func loadData() {
	var err error
	dto.InteractionsData, err = ReadCSV(viper.GetString("data_directory.interactions"))
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, row := range dto.InteractionsData {
		userID, _ := strconv.Atoi(row[0])
		itemID, _ := strconv.Atoi(row[1])
		watchedPct, _ := strconv.ParseFloat(row[4], 64)
		dto.Interactions = append(dto.Interactions, models.Interactions{
			UserID:     userID,
			ItemID:     itemID,
			WatchedPCT: watchedPct,
		})
	}
	dto.UserData, err = ReadCSV(viper.GetString("data_directory.users"))
	if err != nil {
		log.Fatal(err.Error())
	}

	dto.ItemsData, err = ReadCSV(viper.GetString("data_directory.items"))
	if err != nil {
		log.Fatal(err.Error())
	}
	//items
	for _, row := range dto.ItemsData {
		itemID, _ := strconv.Atoi(row[0])
		contentType := row[1]
		title := row[2]
		titleOrig := row[3]
		year, _ := strconv.ParseFloat(row[4], 64)
		genres := SplitData(row[5])
		forKids, _ := strconv.ParseBool(row[7])

		directors := SplitData(row[10]) // Добавляем режиссеров
		actors := SplitData(row[11])    // Добавляем актеров
		dto.Items = append(dto.Items, models.Item{
			ID:          itemID,
			ContentType: contentType,
			Title:       title,
			TitleOrig:   titleOrig,
			ReleaseYear: year,
			Genres:      genres,
			ForKids:     forKids,
			Directors:   directors,
			Actors:      actors,
		})
	}
	//user data
	for _, row := range dto.UserData {
		userID, _ := strconv.Atoi(row[0])
		age := row[1]
		income := row[2]
		sex := row[3]
		kidsFlg, _ := strconv.ParseBool(row[4])
		dto.Users = append(dto.Users, models.User{
			ID:     userID,
			Income: income,
			Sex:    sex,
			Age:    age,

			HasKids: kidsFlg,
		})
	}
	//interactions
	for _, row := range dto.InteractionsData {
		userID, _ := strconv.Atoi(row[0])
		itemID, _ := strconv.Atoi(row[1])
		lastwchd, _ := time.Parse("", row[2])
		watchedPct, _ := strconv.ParseFloat(row[4], 64)
		dto.Interactions = append(dto.Interactions, models.Interactions{
			UserID:      userID,
			ItemID:      itemID,
			LastWatchDT: lastwchd,
			WatchedPCT:  watchedPct,
		})
	}
}

func dbSetup() {
	AppSettings.DBSettings.User = viper.GetString("db.username")
	AppSettings.DBSettings.Database = viper.GetString("db.name")
	AppSettings.DBSettings.Password = viper.GetString("db.password")
	AppSettings.DBSettings.Host = viper.GetString("db.host")
	AppSettings.DBSettings.Port = viper.GetString("db.port")
	AppSettings.DBSettings.SSLMode = viper.GetString("db.sslmode")
}
