package main

import "log"

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {

	var q = []*Node{n}
	array = append(array, n.Name)

	for len(q) > 0 {
		var currNode = q[0]

		for i := 0; i < len(currNode.Children); i++ {
			array = append(array, currNode.Children[i].Name)
			if currNode.Children[i].Children != nil {
				q = append(q, currNode.Children[i])
			}
		}

		q = q[1:]

	}

	return array
}

func main() {
	var headNode = &Node{
		Name: "A",
		Children: []*Node{
			{
				Name: "B",
				Children: []*Node{
					{
						Name:     "E",
						Children: nil,
					},
					{
						Name: "F",
						Children: []*Node{
							{
								Name:     "I",
								Children: nil,
							},
							{
								Name:     "J",
								Children: nil,
							},
						},
					},
				},
			},
			{
				Name:     "C",
				Children: nil,
			},
			{
				Name: "D",
				Children: []*Node{
					{
						Name: "G",
						Children: []*Node{
							{
								Name:     "K",
								Children: nil,
							},
						},
					},
					{
						Name:     "H",
						Children: nil,
					},
				},
			},
		},
	}
	var resultArray = headNode.BreadthFirstSearch([]string{})
	log.Println(resultArray)
}
