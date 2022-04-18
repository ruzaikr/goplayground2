package main

type MinMaxStack struct {
	data []int
	min  []int
	max  []int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{
		data: []int{},
		min:  []int{},
		max:  []int{},
	}
}

func (stack *MinMaxStack) GetSize() int {
	return len(stack.data)
}

func (stack *MinMaxStack) Peek() int {
	return stack.data[len(stack.data)-1]
}

func (stack *MinMaxStack) Pop() int {
	if stack.GetSize() < 1 {
		return -1
	}

	var valueToPop = stack.data[stack.GetSize()-1]

	stack.data = stack.data[:stack.GetSize()-1]

	if valueToPop == stack.GetMin() {
		stack.min = stack.min[:len(stack.min)-1]
	}
	if valueToPop == stack.GetMax() {
		stack.max = stack.max[:len(stack.max)-1]
	}

	return valueToPop
}

func (stack *MinMaxStack) Push(number int) {
	stack.data = append(stack.data, number)
	if number >= stack.GetMax() || len(stack.max) < 1 {
		stack.max = append(stack.max, number)
	}
	if number <= stack.GetMin() || len(stack.min) < 1 {
		stack.min = append(stack.min, number)
	}
}

func (stack *MinMaxStack) GetMin() int {
	if len(stack.min) < 1 {
		return -1
	}
	return stack.min[len(stack.min)-1]
}

func (stack *MinMaxStack) GetMax() int {
	if len(stack.max) < 1 {
		return -1
	}
	return stack.max[len(stack.max)-1]
}

func main() {
	var minMaxStack = NewMinMaxStack()

	//minMaxStack.Push(5)
	//log.Println("minMaxStack.GetMin() returns", minMaxStack.GetMin())
	//log.Println("minMaxStack.GetMax() returns", minMaxStack.GetMax())
	//log.Println("minMaxStack.Peek() returns", minMaxStack.Peek())
	//minMaxStack.Push(7)
	//log.Println("minMaxStack.GetMin() returns", minMaxStack.GetMin())
	//log.Println("minMaxStack.GetMax() returns", minMaxStack.GetMax())
	//log.Println("minMaxStack.Peek() returns", minMaxStack.Peek())
	//minMaxStack.Push(2)
	//log.Println("minMaxStack.GetMin() returns", minMaxStack.GetMin())
	//log.Println("minMaxStack.GetMax() returns", minMaxStack.GetMax())
	//log.Println("minMaxStack.Peek() returns", minMaxStack.Peek())
	//log.Println("minMaxStack.Pop() returns", minMaxStack.Pop())
	//log.Println("minMaxStack.Pop() returns", minMaxStack.Pop())
	//log.Println("minMaxStack.GetMin() returns", minMaxStack.GetMin())
	//log.Println("minMaxStack.GetMax() returns", minMaxStack.GetMax())
	//log.Println("minMaxStack.Peek() returns", minMaxStack.Peek())

	minMaxStack.Push(5)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(5)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(5)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(5)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(8)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(8)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(0)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(8)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(9)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Push(5)
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Pop()
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Pop()
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	minMaxStack.Pop()
	minMaxStack.GetMin()
	minMaxStack.GetMax()
	minMaxStack.Peek()
	//stopped at line 263

}

/*

data:
bottom -> [ 7, 5, 8 ] <- top

min:
[ 7, 5 ]

*/
