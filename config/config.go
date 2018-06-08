package config

import (
	"errors"
	"log"
	"os"

	"github.com/asaskevich/govalidator"
)

type Database struct {
	DB_HOST     string `valid:"required"`
	DB_PORT     string `valid:"required"`
	DB_USER     string `valid:"required"`
	DB_NAME     string `valid:"required"`
	DB_PASSWORD string `valid:"required"`
}

type Webserver struct {
	PORT string `valid:"required"`
}

type Config struct {
	Database  Database
	Webserver Webserver
}

var config *Config

func Init() *Config {
	config = new(Config)

	// Database
	config.Database.DB_HOST = os.Getenv("DB_HOST")
	config.Database.DB_PORT = os.Getenv("DB_PORT")
	config.Database.DB_USER = os.Getenv("DB_USER")
	config.Database.DB_NAME = os.Getenv("DB_NAME")
	config.Database.DB_PASSWORD = os.Getenv("DB_PASSWORD")

	// Webserver
	config.Webserver.PORT = os.Getenv("PORT")

	_, err := govalidator.ValidateStruct(config)
	if err != nil {
		err := errors.New("Invalid Environment values, Error:" + err.Error())
		panic(err)
	}

	log.Println("Config Initialized")
	return config
}

func GetConfig() *Config {
	return config
}
