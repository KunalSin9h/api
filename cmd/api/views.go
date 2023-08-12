package main

import (
	"context"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/kunalsin9h/api/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
views will give views for all blog
This will not update the count

TODO use context with timeout
*/
func (app *App) views(c *fiber.Ctx) error {
	// Get a cursor over all the data
	cursor, err := app.database.Blog.Find(context.TODO(), bson.D{}, options.Find().SetProjection(bson.D{
		{Key: "_id", Value: 0},
	}))

	if err != nil {
		return err
	}

	var response []data.BlogDoc

	if err := cursor.All(context.Background(), &response); err != nil {
		return err
	}

	jsonResponse(c, fiber.StatusOK, true, response)

	return nil
}

/*
blogView will give & update the view of a single
blog.
*/
func (app *App) blogView(c *fiber.Ctx) error {
	// https://docs.gofiber.io/#zero-allocation
	slug := utils.CopyString(c.Params("slug"))

	// get the number of views
	// TODO: set query timeout
	var views data.BlogDoc

	err := app.database.Blog.FindOne(context.TODO(), bson.D{
		{Key: "slug", Value: slug},
	}, options.FindOne().SetProjection(bson.D{
		{Key: "_id", Value: 0},
	})).Decode(&views)

	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Println("No Blog document, creating now...")
	} else if err != nil {
		log.Printf("Failed to get views for slug: %s, the error is: %v", slug, err)
		return err
	}

	// spin an async task to update the view count
	go func() {
		// TODO: set query timeout
		_, err := app.database.Blog.UpdateOne(context.TODO(), bson.D{
			{Key: "slug", Value: slug},
		}, bson.D{
			{Key: "$inc", Value: bson.D{{Key: "views", Value: 1}}},
		}, options.Update().SetUpsert(true))

		if err != nil {
			log.Printf("Failed to update the views for slug: %s and error is: %v", slug, err)
		}
	}()

	// respond with the number of view + 1
	jsonResponse(c, fiber.StatusOK, true, map[string]any{
		// since the views is not updated, we can just +1
		"views": views.Views + 1,
	})
	return nil
}

func jsonResponse(c *fiber.Ctx, code int, success bool, data any) {
	c.Set("Content-Type", "application/json")
	c.Status(code)
	c.JSON(map[string]any{
		"success": success,
		"data":    data,
	})
}
