package main

import "fmt"

type ListNodeII struct {
	Val  int
	Next *ListNodeII
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	list := createListII(arr)
	revList := reverseBetween(list, 2, 4)

	for revList != nil {
		fmt.Println("node1: ", revList.Val)
		revList = revList.Next
	}
}

func reverseBetween(head *ListNodeII, m, n int) *ListNodeII {
	if head == nil {
		return &ListNodeII{}
	}

	dummy := &ListNodeII{Next: head}
	prev := dummy

	for i := 1; i < m; i++ {
		prev = prev.Next
	}

	cur := prev.Next

	for j := m; j < n; j++ {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = prev.Next
		prev.Next = tmp
	}

	return dummy.Next
}

func createListII(arr []int) *ListNodeII {
	n := len(arr)

	if n == 0 {
		return &ListNodeII{}
	}

	var head *ListNodeII

	for i := 0; i < n; i++ {
		node := &ListNodeII{Val: arr[i]}
		node.Next = head
		head = node
	}

	return head
}
