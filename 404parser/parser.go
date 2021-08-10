package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("http://localhost:6888")
	if err != nil {
		log.Fatal(err)
	}

	all_links := getLinks(resp.Body)
	my_links := filterLinks(all_links)

	// Собираем пул, со статусом ответа 404
	link_with_error := codeTesting(my_links)

	// сгенерировать csv со ссылками которые имеют 404
	makeCSV(link_with_error)

}

func makeCSV(links []string) bool {

	make_csv_string := ""

	for _, v := range links {
		make_csv_string = make_csv_string + v + "\n"
	}

	file, err := os.Create("404.csv")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString(make_csv_string)

	return true

}

//Тестируем на 404 и возвращаем отфильтрованный массив
func codeTesting(links []string) []string {

	link_with_error := []string{}

	for _, v := range links {

		right_link := "http://localhost:6888" + v

		resp, err := http.Get(right_link)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == 404 {
			link_with_error = append(link_with_error, v)
		}

	}

	return link_with_error

}

//Функция отфильтровывает слайс, исключая ссылки которые нам не нужны
func filterLinks(links []string) []string {

	link_list := []string{}

	for _, v := range links {

		if strings.Contains(v, ".ru") == false && v != "/" && strings.Contains(v, "/") == true && inSlice(link_list, v) == false {
			link_list = append(link_list, v)
		}
	}

	return link_list

}

// Contains указывает, содержится ли x в a.
func inSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

//Собирает все ссылки
func getLinks(body io.Reader) []string {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:

			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}

		}
	}

}
