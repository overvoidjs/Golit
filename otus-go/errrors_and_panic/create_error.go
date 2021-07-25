package main

import(
  "fmt"
  "errors"
)

func foo() (bool, error) {
  return false, errors.New("Какая то ошибка созданная пакетом errors")
}

func boo() (bool, error) {
  return false, fmt.Errorf("Ошибка %q не найдена", "test")
}

func main() {
  _, err := foo()
  if err != nil {
    fmt.Println(err)
  }

  if _, err := boo(); err != nil {
    fmt.Println(err)
  }

}
