package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var resultListHeadNode *ListNode
	var resultList *ListNode

	for list1 != nil && list2 != nil {
		var newVal int
		if list1.Val <= list2.Val {
			newVal = list1.Val
			list1 = list1.Next
		} else {
			newVal = list2.Val
			list2 = list2.Next
		}

		var listNodeWithNewVal = &ListNode{
			Val:  newVal,
			Next: nil,
		}

		if resultList == nil {
			resultList = listNodeWithNewVal
			resultListHeadNode = resultList
		} else {
			resultList.Next = listNodeWithNewVal
			resultList = resultList.Next
		}
	}

	if list1 == nil && list2 == nil {
		return resultListHeadNode
	}

	var listToAppend *ListNode
	if list1 == nil && list2 != nil {
		listToAppend = list2
	} else if list2 == nil && list1 != nil {
		listToAppend = list1
	}

	if resultList == nil {
		resultList = listToAppend
		resultListHeadNode = resultList
	} else {
		resultList.Next = listToAppend
	}

	return resultListHeadNode
}

func main() {
	var list1 = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	var list2 = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	mergeTwoLists(&list1, &list2)
}
