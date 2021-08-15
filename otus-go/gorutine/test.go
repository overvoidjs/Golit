package main

import (
  "fmt"
)

func repeater(s string, n int) string {
  repeat := ""

  for i := 0; i < n; i++ {
    repeat += s
  }

  ch1 <- "test"
}

func main() {

  ch1 := make(chan string);

  go repeater("s",50)

  for x := range ch1 {
    fmt.Println(x)
  }

}
