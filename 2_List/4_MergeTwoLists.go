package main

import "fmt"

type ListNodeIII struct {
	Val  int
	Next *ListNodeIII
}

func main() {
	arr := []int{5, 3, 1}
	list := createListIII(arr)

	arr1 := []int{6, 4, 2}
	list1 := createListIII(arr1)

	//revList := mergeTwoListsI(list, list1)
	revList := mergeTwoLists(list, list1)

	for revList != nil {
		fmt.Println("node1: ", revList.Val)
		revList = revList.Next
	}
}

func mergeTwoListsI(list1, list2 *ListNodeIII) *ListNodeIII {
	dummyHead := &ListNodeIII{}
	tmp := dummyHead

	if list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			tmp.Next = list2
			list2 = list2.Next
		} else {
			tmp.Next = list1
			list1 = list1.Next
		}
		tmp = tmp.Next
	}
	if list1 != nil {
		tmp.Next = list1
	}
	if list2 != nil {
		tmp.Next = list2
	}

	return dummyHead.Next
}

func mergeTwoLists(list1, list2 *ListNodeIII) *ListNodeIII {
	dummyHead := &ListNodeIII{}
	p := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			p.Next = list2
			list2 = list2.Next
		} else {
			p.Next = list1
			list1 = list1.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}

	return dummyHead.Next
}

func createListIII(arr []int) *ListNodeIII {
	n := len(arr)

	if n == 0 {
		return &ListNodeIII{}
	}

	var head *ListNodeIII

	for i := 0; i < n; i++ {
		node := &ListNodeIII{Val: arr[i]}
		node.Next = head
		head = node
	}

	return head
}
