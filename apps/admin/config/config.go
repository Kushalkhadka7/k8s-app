package config

import (
	"os"
)

type Config struct {
	Server *serverConfig
}

type serverConfig struct {
	Port         string
	Name         string
	Host         string
	AuthUrl      string
	ManagerUrl   string
}

func New() (*Config, error) {
	port:= os.Getenv("PORT")
	name := os.Getenv("NAME")
	host := os.Getenv("HOST")
	authUrl := os.Getenv("AUTH_URL")
	managerUrl := os.Getenv("MANAGER_URL")

	return &Config{
		Server: &serverConfig{
			Port:         port,
			Name:         name,
			Host:         host,
			AuthUrl:      authUrl,
			ManagerUrl:   managerUrl,
		},
	}, nil
}
