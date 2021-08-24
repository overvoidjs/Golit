package main

import (
  "fmt"
  "strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {

  var phone string
  for indx, val := range numbers {
    wow := strconv.FormatUint(uint64(val), 10)
    if indx == 0 {
      phone += "("+wow
    } else if indx == 2 {
      phone += wow+") "
    } else if indx == 6 {
      phone += "-"+wow
    } else {
      phone += wow
    }

  }

  return phone
}

func main() {

  fmt.Println( CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0}) )

}
