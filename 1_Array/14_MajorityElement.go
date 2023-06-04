package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	response := majorityElement(nums)
	fmt.Println(response)
}

func majorityElement(nums []int) int {
	major, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}

		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}
