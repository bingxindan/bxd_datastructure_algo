package main

import "fmt"

func main() {
	digits := []int{1, 2, 9, 0, 9}
	response := plusOne(digits)
	fmt.Println(response)
}

func plusOne(digits []int) []int {
	n := len(digits)

	if n == 0 {
		return nil
	}

	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	// digits中所有元素均为9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
