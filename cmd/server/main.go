package main

import (
	"chat-echo-gorm/internal/api"
	"chat-echo-gorm/internal/model"
	"chat-echo-gorm/internal/store"
	"chat-echo-gorm/internal/ws"
	"log"
	"os"

	"github.com/labstack/echo"
)

func main() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "1234")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_PORT", "5432")

	store.ConnectDatabase()
	store.DB.AutoMigrate(&ws.OutboundMessage{})
	store.DB.AutoMigrate(&model.Message{})
	hub := ws.NewHub()

	go hub.Run()

	e := echo.New()
	api.RegisterRoutes(e, hub)
	log.Println("Server is running on 8080")
	e.Logger.Fatal(e.Start(":8080"))

}
