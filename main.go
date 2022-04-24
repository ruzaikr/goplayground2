package main

import (
	"log"
)

type cell struct {
	combinedVal int
	itemsUsed   []int
}

func KnapsackProblem(items [][]int, capacity int) []interface{} {
	var combinedValsNumRows = len(items)
	var combinedValsNumCols = capacity + 1
	var combinedVals = make([][]cell, combinedValsNumRows)
	var sliceToSlice = make([]cell, combinedValsNumRows*combinedValsNumCols)
	for i := range combinedVals {
		combinedVals[i] = sliceToSlice[:combinedValsNumCols]
		sliceToSlice = sliceToSlice[combinedValsNumCols:]
	}

	for i := 0; i < len(items); i++ {
		var currItemVal = items[i][0]
		var currItemWeight = items[i][1]

		for currCapacity := 0; currCapacity < combinedValsNumCols; currCapacity++ {
			if currCapacity == 0 {
				combinedVals[i][currCapacity].combinedVal = 0
				continue
			}

			if i == 0 {
				if currCapacity >= currItemWeight {
					combinedVals[i][currCapacity] = cell{
						combinedVal: currItemVal,
						itemsUsed:   append(combinedVals[i][currCapacity].itemsUsed, i),
					}
					continue
				}
				combinedVals[i][currCapacity].combinedVal = 0
				continue
			}

			var combinedValUsingCurrItem int

			if currCapacity-currItemWeight >= 0 {
				combinedValUsingCurrItem = combinedVals[i-1][currCapacity-currItemWeight].combinedVal + currItemVal
			}

			var combinedValWithoutUsingCurrItem = combinedVals[i-1][currCapacity].combinedVal
			var itemsUsed = combinedVals[i-1][currCapacity].itemsUsed

			var maxCurrCombinedVal = combinedValWithoutUsingCurrItem

			if combinedValUsingCurrItem > combinedValWithoutUsingCurrItem {
				maxCurrCombinedVal = combinedValUsingCurrItem
				itemsUsed = append(combinedVals[i-1][currCapacity-currItemWeight].itemsUsed, i)
			}

			combinedVals[i][currCapacity] = cell{
				combinedVal: maxCurrCombinedVal,
				itemsUsed:   itemsUsed,
			}
		}
	}

	if combinedVals[combinedValsNumRows-1][combinedValsNumCols-1].itemsUsed == nil {
		combinedVals[combinedValsNumRows-1][combinedValsNumCols-1].itemsUsed = []int{}
	}

	// Replace return below.
	return []interface{}{
		combinedVals[combinedValsNumRows-1][combinedValsNumCols-1].combinedVal, // total value
		combinedVals[combinedValsNumRows-1][combinedValsNumCols-1].itemsUsed,   // item indices
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
