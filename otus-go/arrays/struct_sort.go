package main

import(
  "fmt"
  "strings"
  "sort"
)

type goodList struct{
  Word string
  Count int
}

func main()  {

  analysString := "ехал грека через реку видит грека в реке рак"

  listOfWords := strings.Fields(analysString)

  goodArr := []goodList{}

  //Распаковка и подсчет слов
  for _, val := range listOfWords {

    inStruct := inStruct(val, goodArr)
    if inStruct == -1{
      goodArr = append(goodArr, goodList{Word:val, Count:1})
    } else {
      goodArr[inStruct].Count = goodArr[inStruct].Count + 1
    }
  }


  //SORT
  sort.Slice(goodArr, func(i, j int) bool {

      if goodArr[i].Count != goodArr[j].Count {
          return goodArr[i].Count > goodArr[j].Count
      }

      return goodArr[i].Word > goodArr[j].Word
  })

  fmt.Println(goodArr)

}

func inStruct( findKey string, list []goodList ) int{

  for i := 0;i < len(list);i++ {

    if findKey == list[i].Word{
      return i
    }

  }

  return -1
}
