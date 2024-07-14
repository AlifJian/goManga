package main

import (
	"MangaApi/router"
	"MangaApi/util"

	"github.com/gofiber/fiber/v2"
)

var app fiber.App = *fiber.New()

func main() {
	router.Route(&app)

	port := util.EnvPortOr("3000")

	app.Use("/", func(c *fiber.Ctx) error {
		return c.Status(404).SendString("404 Page Not Found")
	})
	app.Listen(port)
}
