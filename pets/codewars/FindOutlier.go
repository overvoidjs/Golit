package main

import (
  "fmt"
)

func FindOutlier(integers []int) int {

  mod := ""

  for _, value := range integers {

    //Определяем модель принятия решения
    if mod == "" {
      if len(integers) > 2 {
        a, b, c := integers[0], integers[1], integers[len(integers)-1]

        if ( a % 2 == 0 && b % 2 == 0 ) || (a % 2 == 0 && c % 2 == 0) || (b % 2 == 0 && c % 2 == 0) {
          mod = "even"
        } else {
          mod = "odd"
        }
      }
    }

    if mod == "even" && value % 2 != 0 {
      return value
    }

    if mod == "odd" && value % 2 == 0 {
      return value
    }


  }


  return 0

}

func main() {
  fmt.Println( FindOutlier([]int{0,1}) )
}
