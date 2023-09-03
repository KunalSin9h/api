package main

import (
	"encoding/json"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type RequestPayload struct {
	Data []PostMeta `json:"data"`
}

type PostMeta struct {
	Id          string `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Published   string `json:"published"`
	Content     string `json:"content"`
}

func (app *App) addDocuments(c *fiber.Ctx) error {
	index := c.Params("index")

	data := c.Request().Body()

	var jsonData RequestPayload

	if err := json.Unmarshal(data, &jsonData); err != nil {
		slog.Error("Failed to marshal request payload: %v", err.Error())
		return err
	}

	if err := app.meilisearch.AddDocument(index, jsonData.Data); err != nil {
		slog.Error("Failed to add document in meilisearch: %v", err.Error())
		return err
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
