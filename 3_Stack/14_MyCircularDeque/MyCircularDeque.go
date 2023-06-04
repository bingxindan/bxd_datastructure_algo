package main

import "fmt"

func main() {
	obj := Constructor(3)
	response1 := obj.InsertLast(1)
	fmt.Println(response1)
	response2 := obj.InsertLast(2)
	fmt.Println(response2)
	response3 := obj.InsertFront(3)
	fmt.Println(response3)
	response4 := obj.InsertFront(4)
	fmt.Println(response4)
	response5 := obj.GetRear()
	fmt.Println(response5)
	response6 := obj.IsFull()
	fmt.Println(response6)
	response7 := obj.DeleteLast()
	fmt.Println(response7)
	response8 := obj.InsertFront(4)
	fmt.Println(response8)
	response9 := obj.GetFront()
	fmt.Println(response9)
}

type MyCircularDeque struct {
	front, rear int
	elements    []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{elements: make([]int, k+1)}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}

	q.front = (q.front - 1 + len(q.elements)) % len(q.elements)
	q.elements[q.front] = value

	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}

	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)

	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}

	q.front = (q.front + 1) % len(q.elements)

	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.rear = (q.rear - 1 + len(q.elements)) % len(q.elements)
	return true
}

func (q MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[(q.rear-1+len(q.elements))%len(q.elements)]
}

func (q MyCircularDeque) IsEmpty() bool {
	return q.rear == q.front
}

func (q MyCircularDeque) IsFull() bool {
	return (q.rear+1)%len(q.elements) == q.front
}
