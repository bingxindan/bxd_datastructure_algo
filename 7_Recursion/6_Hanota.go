package main

import "fmt"

func main() {

}

func hanota(A, B, C []int) []int {
	if len(A) == 0 {
		return A
	}

	var move func(from, to, other *[]int, n int)
	move = func(from, to, other *[]int, n int) {
		if n == 1 {
			*to = append(*to, (*from)[len(*from)-1])
			*from = (*from)[:len(*from)-1]
			fmt.Println("此时的from = ", from)
			return
		}
		move(from, other, to, n-1)
		*to = append(*to, (*from)[len(*from)-1])
		*from = (*from)[:len(*from)-1]
		move(other, to, from, n-1)
	}
	move(&A, &C, &B, len(A))
	return C
}
