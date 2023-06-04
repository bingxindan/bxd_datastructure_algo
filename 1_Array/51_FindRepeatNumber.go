package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 10000)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		arr[i] = rand.Intn(10)
	}
	fmt.Println("生成的原数组为:", arr)

	res := FindRepeatNumber(arr)

	fmt.Println("res:", res)
}

func FindRepeatNumber(nums []int) map[int]int {
	res := make(map[int]int, 0)

	//特殊判断
	if nums == nil || len(nums) == 0 {
		return res
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return res
		}
	}

	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		if _, exi := res[tmp]; exi {
			res[tmp]++
		} else {
			res[tmp] = 1
		}
	}

	fmt.Println("res:", res)

	res1 := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		tmp1 := nums[i]
		res1[tmp1]++
	}

	fmt.Println("res1:", res1)

	return res
}
