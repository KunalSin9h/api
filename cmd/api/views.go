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

// TODO
func (app *App) views(c *fiber.Ctx) error {
	c.WriteString("TODO: Total views count")
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

	c.Set("Content-Type", "application/json")
	c.Status(fiber.StatusOK)

	c.JSON(map[string]any{
		"success": true,
		"data": map[string]any{
			"views": views.Views + 1,
		},
	})

	return nil
}
