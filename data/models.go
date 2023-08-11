package data

import (
	"context"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	Blog   *mongo.Collection
}

func (mdb *MongoDB) Connect(connString string) error {

	// Exponential Backing
	for i := 1; i <= 5; i++ {

		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString))

		if err != nil {
			if i == 5 {
				return err
			}

			time.Sleep(time.Duration(math.Pow(float64(i), 2)) * time.Second)
		} else {
			mdb.Client = client
			mdb.blog = mdb.Client.Database("api").Collection("blog")

			break
		}

	}

	return nil
}
