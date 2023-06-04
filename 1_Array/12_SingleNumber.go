package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}
	response := singleNumber(nums)
	fmt.Println(response)
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
