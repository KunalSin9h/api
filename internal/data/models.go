package data

import (
	"context"
	"log"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client  *mongo.Client
	Blog    *mongo.Collection
	Timeout time.Duration
}

type BlogDoc struct {
	Slug  string `bson:"slug" json:"slug"`
	Views int64  `bson:"views" json:"views"`
}

func (mdb *MongoDB) Connect(connString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), mdb.Timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))

	if err != nil {
		return err
	}

	max_attempt := 5

	// Exponential Backing
	for i := 1; i <= max_attempt; i++ {
		log.Printf("Trying to ping mongodb to check if we are connected... [%d/%d]\n", i, max_attempt)

		err := client.Ping(context.Background(), readpref.Primary())

		if err != nil {
			if i == max_attempt {
				return err
			}
			time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
		} else {
			mdb.Client = client
			mdb.Blog = mdb.Client.Database("api").Collection("blog")
			log.Printf("Successfully connected to db")
			break
		}
	}

	return nil
}
