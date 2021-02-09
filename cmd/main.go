package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yongDataScince/rest-api"
	"github.com/yongDataScince/rest-api/pkg/handler"
	"github.com/yongDataScince/rest-api/pkg/repository"
	"github.com/yongDataScince/rest-api/pkg/service"
	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConf(); err != nil {
		logrus.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: 		viper.GetString("db.host"),
		Port: 		viper.GetString("db.port"),
		Username: 	viper.GetString("db.username"),
		Password: 	os.Getenv("PASSWORD_DB"),
		DBName: 	viper.GetString("db.dbname"),
		SSLMode: 	viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatal(err)
	}


	rep := repository.NewRepository(db)
	services := service.NewService(rep)
	handlers := handler.NewHandlers(services)
	srv := new(rest.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}

func initConf() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}