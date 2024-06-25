package util

import (
	"MangaApi/model"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Try() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func StrToInt(str string) int {
	out, err := strconv.Atoi(str)

	if err != nil {
		panic("String Converter Error")
	}

	return out
}

func GetLimitIndex(limit string, index string) (int, int) {
	limits := StrToInt(limit)
	indexs := StrToInt(index)

	if limits >= 30 {
		limits = 30
	} else if limits <= 0 {
		limits = 1
	}

	if indexs >= 30 {
		indexs = 0
	}

	return limits, indexs
}

func ScrapeList(doc *goquery.Document, limit int, index int) *[]model.Manga {
	ind := 0
	dataCollection := make([]model.Manga, 0)
	doc.Find("#series-list").Children().Each(func(i int, value *goquery.Selection) {

		isAdult := false
		genre := ""
		value.Find(".item-genre").Children().Each(func(i int, s *goquery.Selection) {
			if s.Text() == "Adult" || s.Text() == "Yaoi(BL)" || s.Text() == "Hentai" || s.Text() == "Mature" {
				isAdult = true
			}
			genre += s.Text() + " , "
		})

		if !isAdult {
			if ind < limit+index && ind >= index {
				row := new(model.Manga)

				row.Title = value.Find(".item-title").Text()

				row.Genre = genre

				url, _ := value.Find(".item-cover").Attr("href")
				row.MangaUrl = "https://wto.to" + url

				_, isIndonesian := value.Find(".flag-indonesia").Attr("class")
				row.Indonesian = isIndonesian

				imgUrl, _ := value.Find(".item-cover").Children().First().Attr("src")
				row.ImageUrl = imgUrl

				row.Chapter = value.Find(".item-volch").Children().First().Text()

				if row.Chapter == "" {
					row.Id = ""
					row.ChapterUrl = ""
					row.Uploader = ""
				} else {
					chapterId, _ := value.Find(".item-volch").Children().First().Attr("href")
					row.Id = strings.Split(chapterId, "chapter/")[1]

					row.ChapterUrl = "https://wto.to" + chapterId

					uploader, uploaderExist := value.Find(".item-volch").Children().Last().Children().First().Attr("href")
					uploadTime := value.Find(".item-volch").Children().Last().Children().Last().Text()
					if !uploaderExist {
						row.Uploader = "Not Found" + uploadTime
					} else {
						row.Uploader = strings.Split(uploader, "/")[3] + uploadTime
					}
				}

				dataCollection = append(dataCollection, *row)
			}
			ind++
		}
	})

	return &dataCollection
}

func EnvPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
