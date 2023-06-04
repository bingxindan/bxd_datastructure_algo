package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
