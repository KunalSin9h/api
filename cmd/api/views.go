package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

/*
getViews will give the views of the blog with slug
this will not update the views
*/
func (app *App) getViews(c *fiber.Ctx) error {
	slug := c.Params("slug")
	blog, err := app.database.GetBlog(slug)

	if err != nil {
		log.Printf("Failed to get BlogDoc for slug: %s, the error is: %v", slug, err)
		return err
	}

	jsonResponse(c, fiber.StatusOK, true, map[string]any{
		"views": blog.Views,
	})
	return nil
}

/*
blogView will give & update the view of a single
blog.
*/
func (app *App) updateViews(c *fiber.Ctx) error {
	// https://docs.gofiber.io/#zero-allocation
	slug := utils.CopyString(c.Params("slug"))

	blog, err := app.database.GetBlog(slug)

	if err != nil {
		log.Printf("Failed to get BlogDoc for slug: %s, the error is: %v", slug, err)
		return err
	}

	// spin an async task to update the view count
	go func() {
		if err := app.database.UpdateBlog(slug); err != nil {
			log.Printf("Failed to update the views for slug: %s", slug)
		}
	}()

	// respond with the number of view + 1
	jsonResponse(c, fiber.StatusOK, true, map[string]any{
		// since the views is not updated, we can just +1
		"views": blog.Views + 1,
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
