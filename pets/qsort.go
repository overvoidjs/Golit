package main

import (
	"fmt"
)

func getHalfOfSlice(arr []int) int {
  x := len(arr) / 2

  if x == 1 {
    x = 2
  }

  return x-1
}

func qsort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else if len(arr) == 2 {
		result_arr := []int{}
		if arr[0] > arr[1] {
			result_arr = append(result_arr, arr[1])
			result_arr = append(result_arr, arr[0])

			return result_arr
		} else {
			return arr
		}
	} else {
    //Вычисляем серидину массива
    half_index := getHalfOfSlice(arr)

    //Записываем наш опорный индекс
    pivot := arr[half_index]

    less_than := []int{}
    more_than := []int{}

    for _, value := range arr {
      if value < pivot {
        less_than = append(less_than, value)
      } else if value > pivot {
        more_than = append(more_than, value)
      }
    }


    less_than = qsort(less_than)
    more_than = qsort(more_than)

    result_arr := append(less_than, pivot)
    result_arr = append(result_arr, more_than...)

    return result_arr

	}

	return arr
}

func main() {
	arr := []int{64,13,22,16,4,5}

	fmt.Println(qsort(arr))
}
