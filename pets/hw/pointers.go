// http://golang-book.ru/chapter-08-pointers.html
package main

import(
  "fmt"
)

func zero(x *int){
  *x = 0
}

func square(x *float64) {
    *x = *x * *x
}

func swap(x *int, y *int){
  tmp := *x
  *x = *y
  *y = tmp
}

func main(){

  // // 1 - Как получить адрес переменной?
  // x := 20
  // fmt.Println(&x)

  // // 2 - Как присвоить значение указателю?
  // x := 200
  // fmt.Println(x)
  // zero(&x)
  // fmt.Println(x)

  // // 3 - Как создать новый указатель?
  // x := new(int)
  // fmt.Println(x)

  // // 4 - Какое будет значение у переменной x после выполнения программы:
  // x := 1.5
  // square(&x)
  // fmt.Println(x)

  // // 5 - Напишите программу, которая меняет местами два числа (x := 1; y := 2; swap(&x, &y) должно дать x=2 и y=1).
  // x := 1
  // y := 2
  // fmt.Println(x,y)
  // swap(&x, &y)
  // fmt.Println(x,y)


}
