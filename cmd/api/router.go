package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (app *App) routes() *fiber.App {

	// Creating an fiber instance
	// TODO Apply fiber configuration from app.config.fiber
	router := fiber.New()

	// Applying cors
	// TODO cors config can be applied from app.config.cors
	router.Use(cors.New())

	// Routes

	/*
	 Home (/)
	 Fow showing list of api as documentation
	*/
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(router.Stack())
	})

	// v1
	v1 := router.Group("v1")

	/*
		Used in https://kunalsin9h.com/blog/slug
		To generate dynamic OG / Twitter Card Images using title
	*/
	v1.Get("/images/:title", app.GenerateImage)

	return router
}
