package main

import (
  "fmt"
  // "strconv"
  "unicode/utf8"
)

func RemoveLastChar(str string) string {
      for len(str) > 0 {
              _, size := utf8.DecodeLastRuneInString(str)
              return str[:len(str)-size]
      }
      return str
}

// func PrevMultOfThree(n int) interface{} {
//
//   if n % 3 == 0 {
//     return n
//   } else {
//     runer := true
//     newString := strconv.Itoa(n)
//
//     for runer {
//
//       if len(newString) == 0 {
//         runer = false
//         return nil
//       }
//
//       nToString := RemoveLastChar(newString)
//       stringToInt, _ := strconv.Atoi(nToString)
//
//       if stringToInt % 3 == 0 {
//         runer = false
//         if stringToInt == 0 {
//           return nil
//         } else {
//           return stringToInt
//         }
//       } else {
//         newString = nToString
//       }
//
//     }
//
//     return nil
//   }
//
// }

func PrevMultOfThree(n int) interface{} {
  for n > 0 {
    if n % 3 == 0 {
      return n
    }
    n /= 10
  }
  return nil
}

func main() {
fmt.Println(PrevMultOfThree(95))


}
