// Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that occur more than once in the input string.
// The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.
package main

import (
"fmt"
"strings"
)

func duplicate_count(s1 string) int {

  s1 = strings.ToUpper(s1)
  stringArr := strings.Split(s1, "")
  myMap := make(map[string]int)

  //Заполняем карту
  for _, value := range stringArr {

    _, ok := myMap[value]
    if !ok {
      myMap[value] = 1
    } else {
      myMap[value]++
    }

  }

  //Узнаем сколько из карты повторяется
  dupli := 0
  for _, value := range myMap {
    if value > 1 {
      dupli++
    }
  }

  return dupli
}

func main() {
  fmt.Println( duplicate_count("abcde") ) // 0
  fmt.Println( duplicate_count("aabbcde") ) // 2
  fmt.Println( duplicate_count("aabBcde") ) //2
  fmt.Println( duplicate_count("indivisibility") ) //1
  fmt.Println( duplicate_count("Indivisibilities") ) //2
  fmt.Println( duplicate_count("aA11") ) //2
  fmt.Println( duplicate_count("ABBA") ) //2
}
