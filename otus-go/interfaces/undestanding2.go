package main

import(
  "fmt"
)

type Logger interface {
  Log()
}

type System struct {
  name string
  htype string
}

func (s System) Log() {
  fmt.Printf("Call %s with type = %s \n", s.name, s.htype)
}

func printLog(log Logger){
  log.Log()
}

func main(){
  command := System{name: "WOW", htype: "test"}

  printLog(command)
}
