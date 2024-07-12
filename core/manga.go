package core

import (
	"MangaApi/model"
	"MangaApi/util"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const site_url = "https://wto.to/"

func GetMangaHome(limit int, index int) *[]model.Manga {

	// Memastikan tidak crash saat panic
	defer util.Try()

	res, err := http.Get("https://mangatoto.org")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("ERROR Code: %d Status : %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	// html, _ := doc.Html()
	// fmt.Println(gohtml.Format(html))

	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	dataCollection := scrapeList(doc, limit, index)
	return dataCollection
}

func GetMangaImg(id string) *[]string {
	// Memastikan tidak crash saat panic
	defer util.Try()

	res, err := http.Get(site_url + "chapter/" + id)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return &[]string{"ERROR:" + res.Status}
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil || doc.Children().Text() == "404 Not Found (1)" {
		return &[]string{"ERROR:501"}
	}

	script := doc.Find("script")

	img := strings.Split(strings.Replace((strings.Split(strings.Split(script.Text(), "imgHttps = [")[1], "];")[0]), "\"", "", -1), ",")

	return &img
}

func SearchManga(title string, limit int, index int) *[]model.Manga {
	url := "https://mangatoto.org/search?word=" + title
	defer util.Try()
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("ERROR Code: %d Status : %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	dataCollection := scrapeList(doc, limit, index)
	return dataCollection
}

func GetMangaSeries(id string) *model.Series {
	defer util.Try()

	res, err := http.Get(site_url + "series/" + id)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}
	var series model.Series

	series.Title = strings.TrimSpace(strings.ReplaceAll(doc.Find(".item-title").Text(), "\n", ""))
	cover, exist := doc.Find(".attr-cover img").Attr("src")
	if exist {
		series.Cover = cover
	}
	series.Alias = strings.TrimSpace(strings.ReplaceAll(doc.Find(".alias-set").Text(), "\n", ""))

	series.Sinopsis = doc.Find(".limit-html").Text()

	doc.Find(".attr-item").Each(func(i int, doc *goquery.Selection) {
		key := doc.Find(".text-muted").Text()
		value := doc.Find("span").First().Text()
		switch key {
		case "Rank:":
			series.Rank = value
		case "Authors:":
			series.Authors = strings.TrimSpace(strings.ReplaceAll(value, "\n", ""))
		case "Genres:":
			genre := ""
			for _, i := range strings.Split(strings.ReplaceAll(value, "\n", ""), ",") {
				genre += strings.TrimSpace(i) + ", "
			}
			series.Genres = genre
		case "Original language:":
			series.OriginLang = value
		case "Translated language:":
			series.TranslatedLang = value
		case "Upload status:":
			series.Status = value
		case "Year of Release:":
			series.Release = value
		}
	})

	chapters := make([]struct {
		Title     string
		ChapterId string
	}, 0)

	doc.Find(".episode-list .item").Each(func(i int, doc *goquery.Selection) {
		row := new(struct {
			Title     string
			ChapterId string
		})
		chapterId, exist := doc.Find("a").Attr("href")
		if exist {
			row.ChapterId = strings.Split(chapterId, "/")[2]
		}
		row.Title = strings.TrimSpace(strings.ReplaceAll(doc.Find("b").Text(), "\n", ""))
		// fmt.Println(row)
		chapters = append(chapters, *row)
	})

	series.Chapter = chapters
	return &series
}

// UTILITIES
func scrapeList(doc *goquery.Document, limit int, index int) *[]model.Manga {
	ind := 0
	dataCollection := make([]model.Manga, 0)
	doc.Find("#series-list").Children().Each(func(i int, value *goquery.Selection) {

		isAdult := false
		genre := ""
		value.Find(".item-genre").Children().Each(func(i int, s *goquery.Selection) {
			if s.Text() == "Adult" || s.Text() == "Yaoi(BL)" || s.Text() == "Hentai" || s.Text() == "Mature" || s.Text() == "Echhi" {
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
				row.SeriesId = strings.Split(url, "/")[2]

				_, isIndonesian := value.Find(".flag-indonesia").Attr("class")
				row.Indonesian = isIndonesian

				imgUrl, _ := value.Find(".item-cover").Children().First().Attr("src")
				row.ImageUrl = imgUrl

				row.Chapter = value.Find(".item-volch").Children().First().Text()

				if row.Chapter == "" {
					row.ChapterId = ""
					row.ChapterUrl = ""
					row.Uploader = ""
				} else {
					chapterId, _ := value.Find(".item-volch").Children().First().Attr("href")
					row.ChapterId = strings.Split(chapterId, "chapter/")[1]

					row.ChapterUrl = "https://wto.to" + chapterId

					uploader, uploaderExist := value.Find(".item-volch").Children().Last().Children().First().Attr("href")
					uploadTime := value.Find(".item-volch").Children().Last().Children().Last().Text()
					if !uploaderExist {
						row.Uploader = "Not Found" + uploadTime
					} else {
						row.Uploader = strings.Split(uploader, "/")[3] + " " + uploadTime
					}
				}

				dataCollection = append(dataCollection, *row)
			}
			ind++
		}
	})

	return &dataCollection
}
