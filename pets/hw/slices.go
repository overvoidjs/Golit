package main

import(
  "fmt"
  "sort"
)

func main(){
  // http://golang-book.ru/chapter-06-arrays-slices-maps.html


  //-----------------------
  // //Чему равна длина среза, созданного таким способом: make([]int, 3, 9)?
  // testArr := make([]int, 3, 9)
  // fmt.Println( len(testArr) ) //3

  //-----------------------
  // // Дан массив:
  // x := [6]string{"a","b","c","d","e","f"}
  // // что вернет вам x[2:5]?
  // fmt.Println(x[2:5]) //c , d, e

  //-----------------------
  // Напишите программу, которая находит самый наименьший элемент в этом списке:
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }

  sort.Ints(x)
  smaller := x[0]

  fmt.Println(smaller)
}
