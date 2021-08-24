package main

import (
  "fmt"
)

func FindOdd(seq []int) int {
  stock := make(map[int]int)

  for _, el := range seq {

    _, ok := stock[el]
    if !ok {
      stock[el] = 1
    } else {
      stock[el]++
    }

  }

  find := 0
  for indx, value := range stock {
    if value % 2 != 0 {
      find = indx
      break
    }
  }

  return find
}

func main() {
 fmt.Println( FindOdd([]int{1,2,2,3,3,3,4,3,3,3,2,2,1}) )
}
