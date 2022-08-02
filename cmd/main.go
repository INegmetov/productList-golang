package main

import (
	"os"

	productApp "github.com/inegmetov/productList-golang"
	"github.com/inegmetov/productList-golang/pkg/handler"
	"github.com/inegmetov/productList-golang/pkg/repository"
	"github.com/inegmetov/productList-golang/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title ProductList App API
// @version 1.0
// @description API Server for ProductList Application

// @host localhost:8080
// @BasePath /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %v", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error loading evn variable: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatal("error to initializing database: %s", err)
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(productApp.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {

		logrus.Fatal("error occurred while running the server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
