package main

func main() {

}

func getDecimalValue(head *ListNode) int {
	re := 0

	for tmp := head; tmp != nil; {
		re = (re << 1) | (tmp.Val)
		tmp = tmp.Next
	}

	return re
}
