package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	list := createListI(arr)

	revList := swapPairs(list)
	for revList != nil {
		fmt.Println("node1: ", revList.Val)
		revList = revList.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	dHead := dummyHead

	for dHead.Next != nil && dHead.Next.Next != nil {
		// 记录下一个节点
		tmp1 := dHead.Next
		// 记录隔一个节点的节点
		tmp2 := dHead.Next.Next

		dHead.Next = tmp2
		tmp1.Next = tmp2.Next
		tmp2.Next = tmp1

		dHead = tmp1
	}

	return dummyHead.Next
}

func createListI(arr []int) *ListNode {
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
