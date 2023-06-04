package main

import (
	"fmt"
	"math"
)

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	/*obj.Push(1)
	obj.Push(2)*/
	response2 := obj.GetMin()
	fmt.Println(response2)
	obj.Pop()
	response := obj.Top()
	response1 := obj.GetMin()
	fmt.Println(response, response1)
}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MinInt64},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, val)
		this.minStack = append(this.minStack, val)
		return
	}

	topVal := this.minStack[len(this.minStack)-1]
	if val < topVal {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, topVal)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
