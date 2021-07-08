package main

import(
  "fmt"
  "math"
  "strconv"
)

func main(){
  var inputs string
  var P float64
  fmt.Println("Введите стартовый капитал")
  fmt.Scanf("%s",&inputs)

  P, _ = strconv.ParseFloat(inputs,2)

  var r float64
  fmt.Println("Введите % ставку в год")
  fmt.Scanf("%s",&inputs)

  r, _ = strconv.ParseFloat(inputs,2)
  r = r / 100

  var n float64 = 1
  var t float64 = 2
  var R1 float64 = 200
  var summ float64 = 0
  var tmp = math.Pow(1+r/n,n*t)


  summ = P * tmp + (R1 * (12.0/n) * (tmp -1)) * (1+r/n) / (r/n)


  fmt.Println(summ)

}
