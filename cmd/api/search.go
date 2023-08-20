package main

import (
	"encoding/json"
	"fmt"

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
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(jsonData)

	if err := app.meilisearch.AddDocument(index, jsonData.Data); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return nil
}

func (app *App) getDocuments(c *fiber.Ctx) error {

	text := c.Query("text")
	index := c.Params("index")

	// TODO: remove this ASAP
	fmt.Println(text)

	data, err := app.meilisearch.SearchDocument(index, text)

	if err != nil {
		return err
	}

	c.JSON(data)

	return nil
}
