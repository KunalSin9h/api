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

	// connect to mongodb client
	if err := app.database.Connect(app.config.mongodbConnString); err != nil {
		log.Fatalf(`Failed to connect to mongodb
		database via connection string %s, error is: %v\n`, app.config, err)
	}

	// Setup up Background and Font required
	// by GenerateImage handler of /image/:title
	// this will allow to fetch these resource once
	// and use multiple times
	if err := app.setUpImageConfig(); err != nil {
		onExitManeuver(app)
		log.Fatalf(`Failed to setup config for /image/:title,
			probably missing assets/font and assets/image folder,
			error is: %v\n`, err,
		)
	}

	log.Printf("Started server on post: %s \n", app.config.applicationPort)

	if err := app.routes().Listen(fmt.Sprintf("%s:%s", app.config.applicationHost, app.config.applicationPort)); err != nil {
		onExitManeuver(app)
		log.Fatalf("Something went wrong, server down with error: %v\n", err)
	}
}

/*
Things which must be done before the app exit
*/
func onExitManeuver(app App) {
	if err := app.database.Client.Disconnect(context.Background()); err != nil {
		log.Printf("Failed to disconnect from mongodb, with error: %v", err)
	}
}
