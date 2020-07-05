package main

import (
	"context"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lmduc/contest-stats/controllers"
	"github.com/lmduc/contest-stats/database"
	"github.com/lmduc/contest-stats/env"
	"github.com/sirupsen/logrus"
)

var (
	db          = initDatabase()
	logger      = logrus.WithFields(logrus.Fields{"package": "main"})
	ctx, cancel = context.WithCancel(context.Background())

	cfgController = controllers.Config{
		Ctx:    ctx,
		Logger: logger,
		DB:     db,
	}
	homeController = controllers.NewHome(&cfgController)
	userController = controllers.NewUser(&cfgController)
)

func main() {
	db.Connect()
	defer db.Close()

	r := gin.Default()
	r.GET("/ping", homeController.HandlePing)
	r.POST("/users", userController.CreateUser)
	r.Run()
}

func initDatabase() database.Database {
	return database.NewPostgresDatabase(
		&database.PostgresConfig{
			Host:     env.GetEnv("DB_HOST"),
			Port:     env.GetEnv("DB_PORT"),
			User:     env.GetEnv("DB_USER"),
			Password: env.GetEnv("DB_PASSWORD"),
			Name:     env.GetEnv("DB_NAME"),
			Ssl:      env.GetEnv("DB_SSL"),
		},
	)
}
