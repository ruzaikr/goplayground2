package main

import (
	"log"
)

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	var waysNoOfRows = len(denoms) + 1
	var waysNoOfCols = n + 1
	var ways = make([][]int, waysNoOfRows)
	var sliceToSlice = make([]int, waysNoOfRows*waysNoOfCols)
	for i := range ways {
		ways[i] = sliceToSlice[:waysNoOfCols]
		sliceToSlice = sliceToSlice[waysNoOfCols:]
	}

	for i := 0; i < waysNoOfRows; i++ {
		for j := 0; j < waysNoOfCols; j++ {
			if j == 0 {
				ways[i][j] = 1
				continue
			}

			if i == 0 {
				ways[i][j] = 0
				continue
			}

			var currDenom int
			var currDenoms = denoms[:i]
			if len(currDenoms) > 0 {
				currDenom = currDenoms[len(currDenoms)-1]
			}

			var noOfWaysUsingCurrDenom int

			if (j - currDenom) >= 0 {
				noOfWaysUsingCurrDenom = ways[i][j-currDenom]
			}

			var noOfWaysWithoutUsingCurrDenom = ways[i-1][j]

			ways[i][j] = noOfWaysUsingCurrDenom + noOfWaysWithoutUsingCurrDenom
		}
	}

	return ways[waysNoOfRows-1][waysNoOfCols-1]
}

func main() {
	var denoms = []int{1, 2, 5}
	var numberOfWaysToMakeChange = NumberOfWaysToMakeChange(5, denoms)
	log.Println(numberOfWaysToMakeChange)
}
