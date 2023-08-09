package main

import (
	"fmt"
	"log"
)

type App struct {
	config      Config
	imageConfig ImageConfig
}

func main() {
	app := App{}

	// Populate all the Config properties
	// using configuration file present at
	// ./config/base.yaml
	if err := app.config.getConfiguration(); err != nil {
		panic(err)
	}

	// Setup up Background and Font required
	// by GenerateImage handler of /image/:title
	// this will allow to fetch these resource once
	// and use multiple times
	if err := app.setUpImageConfig(); err != nil {
		panic(err)
	}

	log.Printf("Started server on post: %s \n", app.config.port)

	log.Fatal(app.routes().Listen(fmt.Sprintf("%s:%s", app.config.host, app.config.port)))
}
