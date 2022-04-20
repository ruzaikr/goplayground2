package main

import (
	"log"
)

func KnapsackProblem(items [][]int, capacity int) []interface{} {
	var combinedValsNumRows = len(items)
	var combinedValsNumCols = capacity + 1
	var combinedVals = make([][]int, combinedValsNumRows)
	var sliceToSlice = make([]int, combinedValsNumRows*combinedValsNumCols)
	for i := range combinedVals {
		combinedVals[i] = sliceToSlice[:combinedValsNumCols]
		sliceToSlice = sliceToSlice[combinedValsNumCols:]
	}

	//var indexesOfItemsUsed []int

	log.Println(combinedVals)
	for i := 0; i < len(items); i++ {
		var currItemVal = items[i][0]
		var currItemWeight = items[i][1]

		for currCapacity := 0; currCapacity < combinedValsNumCols; currCapacity++ {
			if currCapacity == 0 {
				combinedVals[i][currCapacity] = 0
				continue
			}

			if i == 0 {
				if currCapacity >= currItemWeight {
					combinedVals[i][currCapacity] = currItemVal
					continue
				}
				combinedVals[i][currCapacity] = 0
				continue
			}

			var combinedValUsingCurrItem int

			if currCapacity-currItemWeight >= 0 {
				combinedValUsingCurrItem = combinedVals[i-1][currCapacity-currItemWeight] + currItemVal
			}

			var combinedValWithoutUsingCurrItem = combinedVals[i-1][currCapacity]

			var maxCurrCombinedVal = combinedValWithoutUsingCurrItem

			if combinedValUsingCurrItem > combinedValWithoutUsingCurrItem {
				maxCurrCombinedVal = combinedValUsingCurrItem
			}

			combinedVals[i][currCapacity] = maxCurrCombinedVal
		}
	}

	log.Println(combinedVals)

	// Replace return below.
	return []interface{}{
		combinedVals[combinedValsNumRows-1][combinedValsNumCols-1], // total value
		[]int{1, 2}, // item indices
	}
}

func main() {
	var items = [][]int{
		{1, 2},
		{4, 3},
		{5, 6},
		{6, 7},
	}
	var items2 = [][]int{
		{60, 5},
		{50, 3},
		{70, 4},
		{30, 2},
	}
	var answer = KnapsackProblem(items, 10)
	var answer2 = KnapsackProblem(items2, 5)
	log.Println(answer)
	log.Println(answer2)
}
