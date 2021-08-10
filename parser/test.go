package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func postScrape() {
	doc, err := goquery.NewDocument("https://stommarket.ru")
	if err != nil {
		log.Fatal(err)
	}

	// используем CSS селекторы, найденные инспектором в браузере
	// для каждого используем индекс и элемент
	doc.Find(".container .section-title").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		// linkTag := item.Find("a")
		// link, _ := linkTag.Attr("href")
    fmt.Println(title)
		// fmt.Println(link)
	})
}

func main() {
	postScrape()
}
