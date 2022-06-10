package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/jobs?q=go&l&vjk=e13e553fd2aef85c"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	fmt.Println(doc)
	return 0
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed with Status", res.StatusCode)
	}

}
