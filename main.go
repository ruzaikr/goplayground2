package main

import "log"

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) < 1 {
		return -1
	}

	if len(array) == 2 {
		if array[1] > array[0] {
			return array[1]
		}
		return array[0]
	}

	if len(array) == 1 {
		return array[0]
	}

	var i = len(array) - 1
	for i >= 0 {
		var iMinus1 = MaxSubsetSumNoAdjacent(array[:i])
		var iMinus2 = MaxSubsetSumNoAdjacent(array[:i-1])
		var iMinus2PlusI = iMinus2 + array[i]

		i--

		if iMinus1 > iMinus2PlusI {
			return iMinus1
		}
		return iMinus2PlusI
	}
	return -1
}

func main() {
	var mySlice = make([]int, 3, 5)
	mySlice[0] = 4
	mySlice[1] = 6
	mySlice[2] = 1
	var inputArray = []int{75, 105, 120}
	mySlice = mySlice[1 : len(mySlice) : cap(mySlice)-1]
	log.Println(mySlice)
	log.Println(mySlice[:len(mySlice)+2])
	log.Println(MaxSubsetSumNoAdjacent(inputArray))
}
