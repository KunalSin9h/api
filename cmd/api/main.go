package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kunalsin9h/api/data"
)

type App struct {
	database    data.MongoDB
	config      Config
	imageConfig ImageConfig
}

func main() {
	app := App{}

	// Populate all the Config properties
	// using configuration file present at
	// ./config/base.yaml
	app.config.getConfiguration()

	var mongodbConnString = fmt.Sprintf("mongodb://%s:%s@%s:%s",
		app.config.mongodbUsername,
		app.config.mongodbPassword,
		app.config.mongodbHost,
		app.config.mongoPort,
	)

	// connect to mongodb client
	if err := app.database.Connect(mongodbConnString); err != nil {
		log.Fatalf(`Failed to connect to mongodb
		database via connection string %s, error is: %v`, mongodbConnString, err)
	}

	// Disconnect from mongo db database
	// after program finishes
	defer func() {
		if err := app.database.Client.Disconnect(context.Background()); err != nil {
			log.Fatalf(`Failed to disconnect from mongodb client
				with connection string %s, error is: %v`, mongodbConnString, err)
		}
	}()

	// Setup up Background and Font required
	// by GenerateImage handler of /image/:title
	// this will allow to fetch these resource once
	// and use multiple times
	if err := app.setUpImageConfig(); err != nil {
		log.Fatalf(`Failed to setup config for /image/:title,
			probably missing assets/font and assets/image folder,
			error is: %v`, err,
		)
	}

	log.Printf("Started server on post: %s \n", app.config.applicationPort)

	log.Fatal(app.routes().Listen(fmt.Sprintf("%s:%s", app.config.applicationHost, app.config.applicationPort)))
}
