//Разобраться с декодированием Windows-1251 to UTF-8
package main

import (
"fmt"
"log"

"github.com/PuerkitoBio/goquery"
)


func main() {
  testUrl := "https://el-dent.ru/id/filtek-ultimate-flowable-cvet-a2-3930A2-3m.html"
  testPriceQuery := ".cart .floatleft .price"

  fmt.Println("Filtek Ultimate Flowable A2 - Филтек Ультимейт Флоу А2 (3M ESPE)", go_parse(testUrl, testPriceQuery) )

}

func go_parse(target string, priceQuery string) string {


  doc, err := goquery.NewDocument(target)
  if err != nil {
    log.Fatal(err)
  }

  //Определяем кодировку
  dh := doc.Find("head")
  dc := dh.Find("meta[http-equiv]")
  c, _ := dc.Attr("content")
  fmt.Println(c)

  //Парсим
  price := ""

  doc.Find(priceQuery).Each(func(index int, item *goquery.Selection) {
    price += item.Text()
  })



  return price
}
