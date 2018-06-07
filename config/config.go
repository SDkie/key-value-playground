package config

import "os"

type Database struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

type Webserver struct {
	Port string
}

type Config struct {
	Database  Database
	Webserver Webserver
}

var config *Config

func Init() *Config {
	config = new(Config)

	// Database
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.User = os.Getenv("DB_USER")
	config.Database.DbName = os.Getenv("DB_NAME")
	config.Database.Password = os.Getenv("DB_PASSWORD")

	// Webserver
	config.Webserver.Port = os.Getenv("WEBSERVER_PORT")

	return config
}

func GetConfig() *Config {
	return config
}
