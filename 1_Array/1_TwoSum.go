package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	//response := twoSum(nums, target)
	response := hashTwoSum(nums, target)
	fmt.Println(response)
}

func hashTwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	var hashTable = make(map[int]int, 0)

	for k, v := range nums {
		p, ok := hashTable[target-v]
		if ok {
			return []int{p, k}
		}
		hashTable[v] = k
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	for k, v := range nums {
		for j := k + 1; j < len(nums); j++ {
			sum := v + nums[j]
			if sum == target {
				return []int{k, j}
			}
		}
	}
	return nil
}
