package main

import (
	"MangaApi/manga"
	"log"

	"github.com/gofiber/fiber/v2"
)

var app fiber.App = *fiber.New()

func main() {

	app.Get("/", func(c *fiber.Ctx) error {
		home := manga.GetMangaHome()

		return c.JSON(fiber.Map{
			"status": "200",
			"data":   home,
		})
	})

	app.Get("/id/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		imgChapter := manga.GetMangaImg("https://wto.to/chapter/" + id)
		imgLength := len(*imgChapter)

		return c.JSON(fiber.Map{
			"status": "200",
			"data": fiber.Map{
				"imgUrl":    imgChapter,
				"imgLength": imgLength,
			},
		})
	})

	log.Fatal(app.Listen(":3000"))
}
