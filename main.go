package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// https://www.udemy.com/courses/search/?p=1&q=golang&src=ukw
var baseURL string = "https://gall.dcinside.com/board/lists/?id=tree"

func main() {
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageUrl := baseURL + "&page=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

}
func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".bottom_paging_wrap").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with StatusCode :", res.StatusCode)
	}
}
