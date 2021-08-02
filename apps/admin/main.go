package main

import (
	"admin/config"
	"admin/router"
	"admin/server"
	"fmt"
	"os"
)

func main() {

	config, err := config.New()
	if err != nil {
		fmt.Println("Unable to read config from env.")
		panic(err)
	}

	router := router.New()

	server := server.New(config, router)
	if err := server.Start(); err != nil {
		fmt.Println("Unable to start server...%s", err)
		os.Exit(1)
	}

}
