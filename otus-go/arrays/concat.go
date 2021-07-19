package main

import (
	"fmt"
)

func main() {

	zSlice := [][]int{}
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5}
	slice3 := []int{6, 7}

	zSlice = append(zSlice, slice1)
	zSlice = append(zSlice, slice2)
	zSlice = append(zSlice, slice3)

	fmt.Println(zSlice)
	fmt.Println(Concat(zSlice))
}

func Concat(z [][]int) []int {

	concatedSlice := []int{}

	for i := 0; i < len(z); i++ {
		concatedSlice = append(concatedSlice, z[i]...)
	}

	return concatedSlice
}
