package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) {
	dfs(root.Left)

}

func findClosestLeaf(root *TreeNode, k int) int {
	return 0
}

func main() {

	var rootNode = TreeNode{
		Val:   1,
		Left:  &TreeNode{3, nil, nil},
		Right: &TreeNode{2, nil, nil},
	}
	//log.Println("value of closest leaf node is", findClosestLeaf(&rootNode, 2))

	dfs(&rootNode)
}
