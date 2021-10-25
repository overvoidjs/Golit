package main

import (
	"fmt"
)

func qsort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		//Вычисляем серидину массива
		half_index := len(arr) / 2

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
	arr := []int{64, 13}

	fmt.Println(qsort(arr))
}
