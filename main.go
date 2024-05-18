package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/disintegration/imaging"
	"github.com/jung-kurt/gofpdf/v2"
)

type Manga struct {
	Title      string
	Indonesian bool
	Genre      string
	Url        string
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

func Download() []byte {

	fmt.Println("Start Scraping....")
	defer fmt.Println("Download Selesai")
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
			if s.Text() == "Adult" {
				isAdult = true
			}
			genre += s.Text() + " , "
		})

		if !isAdult {
			row := new(Manga)

			row.Title = value.Find(".item-title").Text()

			row.Genre = genre

			url, _ := value.Find(".item-cover").Attr("href")
			row.Url = "https://wto.to" + url

			_, isIndonesian := value.Find(".flag-indonesia").Attr("class")
			row.Indonesian = isIndonesian

			imgUrl, _ := value.Find(".item-cover").Children().First().Attr("src")
			row.ImageUrl = imgUrl

			row.Chapter = value.Find(".item-volch").Children().First().Text()

			chapterId, _ := value.Find(".item-volch").Children().First().Attr("href")
			row.Id = strings.Split(chapterId, "chapter/")[1]

			uploader, _ := value.Find(".item-volch").Children().Last().Children().First().Attr("href")
			uploadTime := value.Find(".item-volch").Children().Last().Children().Last().Text()
			row.Uploader = strings.Split(uploader, "/")[3] + uploadTime

			dataCollection = append(dataCollection, *row)
		}

	})
	dataJson, err := json.MarshalIndent(dataCollection, "", " ")

	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	return dataJson
}

func main() {
	res, err := http.Get("https://wto.to/chapter/2829158")

	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("ERROR Code: %d Status : %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	script := doc.Find("script")
	imgUrl := make([]string, 0)

	img := strings.Split(strings.Replace((strings.Split(strings.Split(script.Text(), "imgHttps = [")[1], "];")[0]), "\"", "", -1), ",")

	fmt.Println(img)
	fmt.Println(imgUrl)

	pdf := gofpdf.New("P", "mm", "A4", "")

	for i, v := range img {
		pdfWidth, pdfHeight := pdf.GetPageSize()

		imgPath := string("img/" + strconv.Itoa(i) + ".jpeg")
		file, err := os.Create(imgPath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		resp, err := http.Get(v)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		fmt.Println(i, " : ", v)
		byte, _ := io.Copy(file, resp.Body)
		fmt.Println("Ukuran File Download : ", byte)

		imgFile, err := imaging.Open(imgPath, imaging.AutoOrientation(true))
		if err != nil {
			log.Fatal(err)
		}
		// var w io.Writer
		// encErr := imaging.Encode(w, imgFile, imaging.JPEG)
		// if encErr != nil {
		// 	log.Fatal(encErr)
		// }
		imgSaveErr := imaging.Save(imgFile, imgPath)
		if imgSaveErr != nil {
			log.Fatal(imgSaveErr)
		}
		pdf.AddPage()

		pdf.ImageOptions(file.Name(), 0, 0, pdfWidth, pdfHeight, false, gofpdf.ImageOptions{ImageType: "JPEG", ReadDpi: true}, 0, "")
	}

	pdferr := pdf.OutputFileAndClose("output.pdf")
	if pdferr != nil {
		log.Fatal(pdferr)
	}

}
