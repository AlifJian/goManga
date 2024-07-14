package router

import (
	"MangaApi/core"
	"MangaApi/model"
	"MangaApi/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	btoto := app.Group("/btoto")

	// Mendapatkan Manga Response Dari Home
	btoto.Get("/", func(c *fiber.Ctx) error {
		limit, index := util.GetLimitIndex(c.Query("limit", "10"), c.Query("index", "0"))

		data := core.GetMangaHome(limit, index)

		return c.JSON(model.ResponseManga{
			Status:  200,
			Message: "OK",
			Data:    *data,
		})
	})

	// Mendapatkan Manga Series Dari SeriesId
	btoto.Get("/series/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		data := core.GetMangaSeries(id)

		return c.JSON(model.ResponseSeries{
			Status:  200,
			Message: "OK",
			Data:    *data,
		})
	})

	// Mendapatkan Image berdasarkan ID Chapter
	btoto.Get("/chapter/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		imgChapter := core.GetMangaImg(id)
		imgLength := len(*imgChapter)

		if strings.Split((*imgChapter)[0], ":")[0] == "ERROR" {
			// string to int
			code := util.StrToInt(strings.Split((*imgChapter)[0], ":")[1])

			return c.Status(code).JSON(fiber.Map{
				"Status":  code,
				"Message": "Internal Server Error",
				"Data":    fiber.Map{},
			})
		}
		return c.JSON(fiber.Map{
			"Status":  "200",
			"Message": "OK",
			"Data": fiber.Map{
				"imgUrl":    imgChapter,
				"imgLength": imgLength,
			},
		})
	})

	// Mendapatkan Manga dengan query
	btoto.Get("/search/", func(c *fiber.Ctx) error {
		limit, index := util.GetLimitIndex(c.Query("limit", "10"), c.Query("index", "0"))
		title := c.Query("title")

		data := core.SearchManga(title, limit, index)

		return c.JSON(model.ResponseManga{
			Status:  200,
			Message: "OK",
			Data:    *data,
		})
	})
}
