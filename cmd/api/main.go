package main

import (
	"log"
	"sample/pkg/config"
	"sample/pkg/di_in"
	// "sample/pkg/di_in"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("can't start server", err)
	}

	server, err := di_in.InitializeAPI(config)
	if err != nil {
		log.Fatal("Can't start server", err)
	} else {
		server.Start()
	}

}
