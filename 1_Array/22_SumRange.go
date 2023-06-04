package main

import "fmt"

func main() {
	i, j := 1, 1
	nums := []int{1}
	obj := ConstructorI(nums)
	response := obj.SumRange(i, j)
	fmt.Println(response)
}

type NumArray struct {
	sums []int
}

func ConstructorI(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums: sums}
}

func (na *NumArray) SumRange(i, j int) int {
	return na.sums[j+1] - na.sums[i]
}
