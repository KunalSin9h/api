package main

import (
	"encoding/json"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type RequestPayload struct {
	Data any `json:"data"`
}

func (app *App) addDocuments(c *fiber.Ctx) error {
	index := c.Params("index")

	data := c.Request().Body()

	var jsonData RequestPayload

	if err := json.Unmarshal(data, &jsonData); err != nil {
		slog.Error("Failed to marshal request payload: %v", err.Error())
		return err
	}

	if err := app.meilisearch.AddDocument(index, jsonData); err != nil {
		slog.Error("Failed to add document in meilisearch: %v", err.Error())
		return nil
	}

	return nil
}

func (app *App) getDocuments(c *fiber.Ctx) error {
	text := c.Query("text")
	index := c.Params("index")

	data, err := app.meilisearch.SearchDocument(index, text)

	if err != nil {
		return err
	}

	c.JSON(data)

	return nil
}
