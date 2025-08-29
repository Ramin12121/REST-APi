package main

import (
	Server "Subscription/pkg/Server"
	"Subscription/pkg/db"
	"Subscription/pkg/handlers"

	_ "Subscription/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Subscription API
// @version 1.0
// @description This is a sample API for managing subscriptions.
// @host localhost:8080
// @BasePath /
func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	database, err := db.InitDB()
	if err != nil {
		logrus.Fatal("Could not connect to DB:")
	}
	e := echo.New()

	Repo := Server.NewRepository(database)
	Service := Server.NewService(Repo)
	Handlers := handlers.NewHandler(Service)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/subscriptions", Handlers.GetAll)
	e.GET("/subscriptionsByFilter", Handlers.GetByFilter)
	e.POST("/subscriptions", Handlers.Post)
	e.PATCH("/subscriptions/:id", Handlers.Patch)
	e.DELETE("/subscriptions/:id", Handlers.Delete)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
