package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	recomendation "recommendation"
	"recommendation/db"
	"recommendation/internal/handler"
	"recommendation/internal/services"
	"recommendation/utils"
	"syscall"
)

func main() {
	utils.ReadSettings()

	db.StartDbConnection()
	_db := db.GetDBConn()

	defer db.DisconnectDB(_db)
	srv := new(recomendation.Server)
	service := services.NewServices(nil)
	handlers := handler.NewHandler(service)

	go func() {
		if err := srv.Run(":8080", handlers.InitRoutes()); err != nil {
			log.Fatal(err.Error())
		}
	}()

	fmt.Println("App started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println(err.Error())
	}
}
