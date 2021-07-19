package main

import(
  "fmt"
  "time"
  "strings"
  "os/exec"
  "os"
)

func main(){

  startString := "......"
  foxRunner := "."
  foxChanger := "!"

  for {

    clearTerminal()

    if strings.ContainsAny(startString, foxRunner) {

      out := strings.Replace(startString, foxRunner, foxChanger, 1)
      startString = out

      fmt.Println(startString)

      time.Sleep(500 * time.Millisecond)


    } else {
      foxTmp := foxRunner
      foxRunner = foxChanger
      foxChanger = foxTmp
    }

  }

}

func clearTerminal(){
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}
