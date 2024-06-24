package core

import (
	"MangaApi/model"
	"MangaApi/util"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetMangaHome(limit int, index int) *[]model.Manga {

	// Memastikan tidak crash saat panic
	defer util.Try()

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
	dataCollection := util.ScrapeList(doc, limit, index)
	return dataCollection
}

func GetMangaImg(url string) *[]string {
	// Memastikan tidak crash saat panic
	defer util.Try()

	res, err := http.Get(url)

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
	dataCollection := util.ScrapeList(doc, limit, index)
	return dataCollection
}
