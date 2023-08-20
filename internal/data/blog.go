package data

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
GetBlog will give the BlogDoc upon a slug
*/
func (mdb *MongoDB) GetBlog(slug string) (BlogDoc, error) {
	var blog BlogDoc

	ctx, cancel := context.WithTimeout(context.Background(), mdb.Timeout)
	defer cancel()

	err := mdb.Blog.FindOne(ctx, bson.D{
		{Key: "slug", Value: slug},
	}, options.FindOne().SetProjection(bson.D{
		{Key: "_id", Value: 0},
	})).Decode(&blog)

	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Println("No Blog document, creating now...")
	} else if err != nil {
		return BlogDoc{}, err
	}

	return blog, nil
}

/*
UpdateBlog will update the blog
*/
func (mdb *MongoDB) UpdateBlog(slug string) error {
	ctx, cancel := context.WithTimeout(context.Background(), mdb.Timeout)
	defer cancel()

	_, err := mdb.Blog.UpdateOne(ctx, bson.D{
		{Key: "slug", Value: slug},
	}, bson.D{
		{Key: "$inc", Value: bson.D{{Key: "views", Value: 1}}},
	}, options.Update().SetUpsert(true))

	if err != nil {
		return err
	}

	return nil
}
