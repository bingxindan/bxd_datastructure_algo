package main

import "fmt"

type Element interface{} //可存入任何类型

type StackTest struct {
	list []Element
}

//初始化栈
func NewStack() *StackTest {
	return &StackTest{
		list: make([]Element, 0),
	}
}

func (s *StackTest) Len() int {
	return len(s.list)
}

//判断栈是否空
func (s *StackTest) IsEmpty() bool {
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}

//入栈
func (s *StackTest) Push(x interface{}) {
	s.list = append(s.list, x)
}

//连续传入
func (s *StackTest) PushList(x []Element) {
	s.list = append(s.list, x...)
}

//出栈
func (s *StackTest) Pop() Element {
	if len(s.list) <= 0 {
		fmt.Println("Stack is Empty")
		return nil
	} else {
		ret := s.list[len(s.list)-1]
		s.list = s.list[:len(s.list)-1]
		return ret
	}
}

//返回栈顶元素，空栈返nil
func (s *StackTest) Top() Element {
	if s.IsEmpty() == true {
		fmt.Println("Stack is Empty")
		return nil
	} else {
		return s.list[len(s.list)-1]
	}
}

//清空栈
func (s *StackTest) Clear() bool {
	if len(s.list) == 0 {
		return true
	}
	for i := 0; i < s.Len(); i++ {
		s.list[i] = nil
	}
	s.list = make([]Element, 0)
	return true
}

//打印测试
func (s *StackTest) Show() {
	len := len(s.list)
	for i := 0; i != len; i++ {
		fmt.Println(s.Pop())
	}
}

type Queue struct {
	stack1, stack2 *StackTest
}

func Construct() Queue {
	return Queue{
		stack1: NewStack(),
		stack2: NewStack(),
	}
}

func (s *Queue) AppendTail(val int) {
	s.stack1.Push(val)
}

func (s *Queue) DeleteHead() int {
	if s.stack2.Len() == 0 {
		if s.stack1.Len() > 0 {
			s.stack2.Push(s.stack1.Pop())
		}
	}

	if s.stack2.Len() > 0 {
		ret := s.stack2.Pop()
		return ret.(int)
	}

	return -1
}

func main() {
	obj := Construct()

	var ele = []int{2, 23, 3453, 564, 5, 67, 4, 123, 66, 57, 34, 809, 8765, 4243, 5765}

	for k, v := range ele {
		fmt.Println(33, k, 44, v)
		obj.AppendTail(v)
	}

	fmt.Println(11, obj.stack1)

	for {
		param := obj.DeleteHead()

		fmt.Println(22, param)

		if param == -1 {
			break
		}
	}
}
