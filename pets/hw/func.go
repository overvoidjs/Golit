// http://golang-book.ru/chapter-07-functions.html
package main

import(
  "fmt"
)

// Функция sum принимает срез чисел и складывает их вместе. Как бы выглядела сигнатура этой функции?
func sum(param []int) int {

  var summer int = 0

  for _, value := range param {
    summer += value
  }

  return summer
}

// Напишите функцию, которая принимает число, делит его пополам и возвращает true в случае, если исходное число чётное, и false, если нечетное.
func isEven(val int) (int, bool) {

  cropVal := val / 2

  if val % 2 == 0 {
    return cropVal, true
  } else {
    return cropVal, false
  }

}

// Напишите функцию с переменным числом параметров, которая находит наибольшее число в списке.
func biggest(args ...int) int {
  bigger := 0

  for _, value := range args {

    if value > bigger {
      bigger = value
    }

  }

  return bigger
}

// Используя в качестве примера функцию makeEvenGenerator напишите makeOddGenerator, генерирующую нечётные числа.
func makeOddGenerator() func() uint {
  i := uint(1)

  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

// Последовательность чисел Фибоначчи определяется как fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Напишите рекурсивную функцию, находящую fib(n).
func fib(x int) int {

  if x <= 1 {
    return x
  } else {
    return fib(x - 1) + fib(x - 2)
  }

}


func main(){
  // // 1
  // xList := []int{1,2,3,4,5}
  // fmt.Println( sum(xList) )

  // // 2
  // fmt.Println( isEven(2) )

  // // 3
  // fmt.Println( biggest(1,99,3) )

  // // 4
  // nextOdd := makeOddGenerator()
  // fmt.Println( nextOdd() )
  // fmt.Println( nextOdd() )
  // fmt.Println( nextOdd() )
  // fmt.Println( nextOdd() )

  // 5
  fmt.Println( fib(15) )
}
