package main

func main() {

}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head

	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
