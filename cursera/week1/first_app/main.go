package main

import (
  "os"
  "fmt"
  "bufio"
  "io"
)

func uniq(input io.Reader, output io.Writer) error {
  in := bufio.NewScanner(input)

  var prev string

  for in.Scan(){
    txt :=in.Text()

    if prev == txt {
        continue
    }

    if txt < prev {
      return fmt.Errorf("file not sorted")
    }

    fmt.Fprintln(output, txt)
    prev = txt
  }

  return nil
}

func main()  {
  err := uniq(os.Stdin, os.Stdout)
  if err != nil {
    panic(err.Error())
  }
}
