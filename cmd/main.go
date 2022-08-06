package main

import (
	"github.com/bogdansemkin/image-softcery/configs"
	"github.com/bogdansemkin/image-softcery/pkg/handlers"
	"github.com/bogdansemkin/image-softcery/pkg/repositories"
	"github.com/bogdansemkin/image-softcery/pkg/services"
	"github.com/bogdansemkin/image-softcery/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// @title   Image App API
// @version 1.0
// @description API Server for upload-download image

// @host     localhost:8000
// @BasePath /
func main() {
	srv := server.Server{}
	config := configs.Configs{}

	if err := config.InitConfig(); err != nil {
		logrus.Errorf("Error during init configs, %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error during init .env, %s", err)
	}

	db, err := repositories.NewPostgresDB(repositories.Config{
		Username: viper.GetString("db.username"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBname:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Errorf("Error during loading configs to db, %s", err)
	}

	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)
	srv.Run("8000", handler.InitRoutes())
}
