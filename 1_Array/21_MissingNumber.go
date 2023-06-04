package main

import "fmt"

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	response := missingNumber(nums)
	fmt.Println(response)
}

func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}

	return total - arrSum
}
