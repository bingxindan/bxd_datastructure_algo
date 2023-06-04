package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	//list := createList(arr)
	list := createListI(arr)
	/*for list != nil {
		fmt.Println("node: ", list.Val)
		list = list.Next
	}*/

	revList := reverseListIII(list)
	for revList != nil {
		fmt.Println("node1: ", revList.Val)
		revList = revList.Next
	}
}

func reverseListIII(head *ListNode) *ListNode {
	if head == nil {
		return &ListNode{}
	}

	var tail *ListNode

	for head != nil {
		// 临时保存下一个链表的节点，等着后面重新返转使用
		tmp := head.Next
		// 将指针指向前一个节点
		head.Next = tail
		// 保存新链表的头节点
		tail = head
		// 保证链表的循环
		head = tmp
	}

	return tail
}

func reverseListII(head *ListNode) *ListNode {
	if head == nil {
		return &ListNode{}
	}

	var prev *ListNode

	for head != nil {
		// 临时保存下一个节点，避免找不到
		tmp := head.Next
		// 将节点指向前一个节点
		head.Next = prev
		prev = head
		// 走向下一个节点
		head = tmp
	}

	return prev
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

// createList 生成链表
// @param
// @author zhangming
// @date 2023/5/24-13:53
func createList(arr []int) *ListNode {
	length := len(arr)

	if length == 0 {
		return nil
	}

	head := &ListNode{}

	for i := 0; i < length; i++ {
		val := arr[i]
		node := &ListNode{Val: val}
		node.Next = head
		head = node
	}

	return head
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var tail *ListNode

	for head != nil {
		node := head.Next
		head.Next = tail
		tail = head

		head = node
	}

	return tail
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
