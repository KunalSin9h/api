package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"github.com/gofiber/fiber/v2"
)

/*
ImageConfig holds the resources
such as background image
*/
type ImageConfig struct {
	backgroundImage image.Image
}

/*
GenerateImage is an handler to generate an image with text provided
by request
*/
func (app *App) GenerateImage(c *fiber.Ctx) error {
	// Title to put in the image
	title := c.Params("title")

	title = strings.ReplaceAll(title, "%20", " ")

	// Creating a new GG Context with dimension 1200 * 630
	dc := gg.NewContext(1200, 630)

	dc.DrawImage(app.imageConfig.backgroundImage, 0, 0)

	// TOTO check why the underline implementation of LoadFontFace
	// is not thread safe
	if err := dc.LoadFontFace("assets/fonts/RobotoSlab.ttf", 80); err != nil {
		return err
	}

	textRightMargin := 60.0
	textTopMargin := 90.0

	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(dc.Width()) - (2.0 * textRightMargin)

	dc.SetColor(color.White)

	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.6, gg.AlignLeft)

	// Buffer to store new image bytes
	var imageBuff bytes.Buffer

	// Encoding new image into imageBuff
	if err := jpeg.Encode(&imageBuff, dc.Image(), nil); err != nil {
		return err
	}

	c.Set("Content-type", "image/jpeg")
	c.Status(200)

	// Send the new image data
	c.Write(imageBuff.Bytes())

	return nil
}

/*
setUpImageConfig to populate ImageConf
*/
func (app *App) setUpImageConfig() error {
	file, err := os.Open("assets/images/og-templ.jpg")
	if err != nil {
		return nil
	}
	defer file.Close()

	app.imageConfig.backgroundImage, err = jpeg.Decode(file)

	if err != nil {
		return err
	}

	return nil
}
