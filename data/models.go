package data

import (
	"context"
	"log"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	Blog   *mongo.Collection
}

type BlogDoc struct {
	Slug  string `bson:"slug" json="slug"`
	Views int64  `bson:"views" json="views"`
}

func (mdb *MongoDB) Connect(connString string) error {

	// Exponential Backing
	for i := 1; i <= 5; i++ {

		log.Printf("Trying to connect to mongodb... [%d/%d]\n", i, 5)

		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString))

		if err != nil {
			if i == 5 {
				return err
			}

			time.Sleep(time.Duration(math.Pow(float64(i), 2)) * time.Second)
		} else {
			mdb.Client = client
			mdb.Blog = mdb.Client.Database("api").Collection("blog")
			log.Printf("Successfully connected to db")
			break
		}

	}

	return nil
}