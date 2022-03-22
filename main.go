package main

import (
	"log"
	"sort"
)

func SortedSquaredArray(array []int) []int {
	if len(array) < 1 {
		return []int{}
	}
	var resultSlice = make([]int, len(array))
	for i := 0; i < len(array); i++ {
		resultSlice[i] = array[i] * array[i]
	}
	sort.Ints(resultSlice)
	return resultSlice
}

func main() {
	var inputSlice = []int{-2, -1}
	var squaredSlice = SortedSquaredArray(inputSlice)
	log.Println(squaredSlice)
}
