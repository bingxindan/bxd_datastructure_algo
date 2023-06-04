package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//arr := []int{5, 4, 3, 2, 1}
	arr := []int{1}
	list := createList(arr)

	reorderList(list)

	for list != nil {
		fmt.Println("node: ", list.Val)
		list = list.Next
	}
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 求中间节点
	middle := middleNode(head)

	// 求出后半段链表
	tail := middle.Next
	// 只留出前增段链表
	middle.Next = nil
	// 前半段链表
	prev := head

	// 从中间后的节点开始翻转链表
	tail = reverseList(tail)

	// 将前半段链表和后半段链表翻转后的链表合并
	mergeList(prev, tail)
}

func mergeList(prev, tail *ListNode) {
	var prevTmp, tailTmp *ListNode
	for prev != nil && tail != nil {
		prevTmp = prev.Next
		tailTmp = tail.Next

		prev.Next = tail
		prev = prevTmp

		tail.Next = prev
		tail = tailTmp
	}
}

func reverseList(head *ListNode) *ListNode {
	var tail *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = tail
		tail = head
		head = tmp
	}

	return tail
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
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
