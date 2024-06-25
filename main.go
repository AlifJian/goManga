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
	app.Listen(port)
}
