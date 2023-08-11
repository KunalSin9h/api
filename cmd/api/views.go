package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (app *App) views(c *fiber.Ctx) error {
	return nil
}

/*
blogView will give & update the view of a single
blog.
*/
func (app *App) blogView(c *fiber.Ctx) error {
	slug := c.Params("slug")

	// get the number of views
	// TODO: set query timeout
	views := app.database.Blog.FindOne(context.TODO(), bson.D{
		{Key: "slug", Value: slug},
	}, options.FindOne().SetProjection(bson.D{
		{Key: "_id", Value: 0},
		{Key: "slug", Value: 0},
		{Key: "views", Value: 1},
	}))

	/**
	ISSUE: fiber uses same variable, so we can use slug in async
	when the handler finished
	*/

	// spin an async task to update the view count
	go func() {
		// TODO: set query timeout
		app.database.Blog.UpdateOne(context.TODO(), bson.D{
			{Key: "slug", Value: slug},
		}, bson.D{
			{Key: "$inc", Value: "views"},
		})

	}()

	// respond with the number of view + 1

	return nil
}
