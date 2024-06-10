package manga

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Manga struct {
	Title      string
	Indonesian bool
	Genre      string
	MangaUrl   string
	ChapterUrl string
	ImageUrl   string
	Id         string
	Chapter    string
	Uploader   string
}

func wait(seconds int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Slept ", seconds, " seconds ..")
}

func GetMangaHome() *[]Manga {

	fmt.Println("Scraping....")
	defer fmt.Println("Scraping Selesai")
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

	dataCollection := make([]Manga, 0)

	doc.Find("#series-list").Children().Each(func(i int, value *goquery.Selection) {
		isAdult := false
		genre := ""
		value.Find(".item-genre").Children().Each(func(i int, s *goquery.Selection) {
			if s.Text() == "Adult" || s.Text() == "Yaoi(BL)" {
				isAdult = true
			}
			genre += s.Text() + " , "
		})

		if !isAdult {
			row := new(Manga)

			row.Title = value.Find(".item-title").Text()

			row.Genre = genre

			url, _ := value.Find(".item-cover").Attr("href")
			row.MangaUrl = "https://wto.to" + url

			_, isIndonesian := value.Find(".flag-indonesia").Attr("class")
			row.Indonesian = isIndonesian

			imgUrl, _ := value.Find(".item-cover").Children().First().Attr("src")
			row.ImageUrl = imgUrl

			row.Chapter = value.Find(".item-volch").Children().First().Text()

			chapterId, _ := value.Find(".item-volch").Children().First().Attr("href")
			row.Id = strings.Split(chapterId, "chapter/")[1]

			row.ChapterUrl = "https://wto.to" + chapterId

			uploader, _ := value.Find(".item-volch").Children().Last().Children().First().Attr("href")
			uploadTime := value.Find(".item-volch").Children().Last().Children().Last().Text()
			row.Uploader = strings.Split(uploader, "/")[3] + uploadTime

			dataCollection = append(dataCollection, *row)
		}

	})
	return &dataCollection
}

func GetMangaImg(url string) *[]string {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("ERROR Code: %d Status : %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil || doc.Children().Text() == "404 Not Found (1)" {
		panic("status: 501")
	}

	script := doc.Find("script")

	img := strings.Split(strings.Replace((strings.Split(strings.Split(script.Text(), "imgHttps = [")[1], "];")[0]), "\"", "", -1), ",")

	return &img
}
