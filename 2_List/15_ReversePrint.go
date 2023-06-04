package main

func main() {

}

func reversePrint(head *ListNode) []int {
	st := make([]int, 0)
	ans := make([]int, 0)

	for it := head; it != nil; it = it.Next {
		st = append(st, it.Val)
	}

	for i := len(st) - 1; i >= 0; i-- {
		ans = append(ans, st[i])
	}

	return ans
}
