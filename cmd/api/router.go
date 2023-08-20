package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (app *App) routes() *fiber.App {
	// Creating an fiber instance
	// TODO Apply fiber configuration from app.config.fiber
	router := fiber.New()

	// Applying cors
	// TODO cors config can be applied from app.config.cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: "http://*, https://*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
		AllowCredentials: true,
		ExposeHeaders:    "Link",
		MaxAge:           300,
	}))

	router.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Routes

	// Home (/)
	// Fow showing list of api as documentation
	router.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(router.Stack())
	})

	// v1
	v1 := router.Group("v1")

	// Used in https://kunalsin9h.com/blog/slug
	// To generate dynamic OG / Twitter Card Images using title
	v1.Get("/image/:title", app.GenerateImage)

	// Used in https://kunalsin9h.com/blog
	// To show total view on the blog post
	v1.Get("/views/:slug", app.getViews)

	// Update & Give the updated view count
	// for the blog  with slug
	v1.Post("/views/:slug", app.updateViews)

	// Get documents which will be return from meilisearch
	// when provided with ?text="search text"
	v1.Get("/search/:index", app.getDocuments)

	// Add documents
	v1.Post("/search/:index", app.addDocuments)

	return router
}
