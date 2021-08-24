package main

import (
  "fmt"
  "strings"
)


//Функция разворачивает строку
func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return
}

func SpinWords(str string) string {

  strToArr := strings.Split(str, " ")
  strToReturn := ""

  if len(strToArr) <= 0 {
    panic("К успеху шел")
  }

  for i, slovo := range strToArr {
    if len(slovo) >= 5 {
      strToReturn += Reverse(slovo)
    } else {
      strToReturn += slovo
    }

    if i != len(strToArr) - 1{
      strToReturn += " "
    }
  }

  return strToReturn
}// SpinWords

func main() {
  fmt.Println( SpinWords("Hey fellow warriors") )
}
