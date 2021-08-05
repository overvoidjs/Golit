package main

import (
  "container/list"
  "fmt"
)

func main() {
    // Создаем новый список
    // и помещаем в него несколько чисел.
    l := list.New()
    e4 := l.PushBack(4)
    e1 := l.PushFront(1)
    l.InsertBefore(3, e4)
    l.InsertAfter(2, e1)
    fmt.Printf("Длина списка %v\n", l.Len())

    fmt.Println("Итерируем по списку с начала.")
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

    fmt.Println("Итерируем по списку с конца.")
    for e := l.Back(); e != nil; e = e.Prev() {
        fmt.Println(e.Value)
    }

}
