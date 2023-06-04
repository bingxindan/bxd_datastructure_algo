package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{6, 5, 4, 3, 2, 1}
	list := createList(arr)
	/*for list != nil {
		fmt.Println("node: ", list.Val)
		list = list.Next
	}*/

	//middle := middleNodeI(list)
	middle := middleNodeII(list)
	fmt.Println("aaaaaa", middle.Val)
}

func middleNodeII(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func middleNodeI(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func createList(arr []int) *ListNode {
	n := len(arr)

	if n == 0 {
		return &ListNode{}
	}

	var head *ListNode

	for i := 0; i < n; i++ {
		node := &ListNode{Val: arr[i]}
		node.Next = head
		head = node
	}

	return head
}
