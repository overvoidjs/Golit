package main

import (
  "fmt"
  "strconv"
)

func DigitalRoot(n int) int {
  newString := strconv.Itoa(n)

  if len(newString) == 1 {
    return n
  } else {

    runner := true
    summer := 0
    for runner {

      for _, val := range newString {
        strToInt, _ := strconv.Atoi(string(val))
        summer += strToInt
      }

      answer := strconv.Itoa(summer)

      if len(answer) == 1 {

        runner = false
        return summer
      } else {
        newString = answer
        summer = 0
      }

    }

  }

  return 0
}

func main() {
  fmt.Println( DigitalRoot(493193) )
}
