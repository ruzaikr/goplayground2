package main

import "log"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func makePositiveIfNegative(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func (tree *BST) FindClosestValue(target int) int {
	var currentNode = tree
	var lowestDiff = 0

	var answer int

	for currentNode != nil {
		var newDiff = makePositiveIfNegative(currentNode.Value - target)

		if newDiff == 0 {
			return target
		}

		if newDiff < lowestDiff || lowestDiff == 0 {
			lowestDiff = newDiff
			answer = currentNode.Value
		}

		if currentNode.Value < target {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}

	}

	return answer
}

func main() {
	var t = &BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &BST{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
				Left:  nil,
				Right: &BST{
					Value: 14,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BST{
				Value: 22,
				Left:  nil,
				Right: nil,
			},
		},
	}
	var closest = t.FindClosestValue(12)
	log.Println(closest)
}
