package main

import (
	"MangaApi/router"

	"github.com/gofiber/fiber/v2"
)

var app fiber.App = *fiber.New()

func main() {
	router.Route(&app)
	app.Listen(":3000")
}
