package main

import (
	"fmt"
	"math"
)

func main() {
	number := 1
	value := 2
	obj := Constructor()
	obj.Add(number)
	response := obj.Find(value)
	fmt.Println(response)
}

type TwoSum struct {
	val map[int]int
	max int
	min int
}

func Constructor() TwoSum {
	val := make(map[int]int)
	return TwoSum{val: val, min: math.MaxInt64, max: math.MinInt64}
}

func (this *TwoSum) Add(number int) {
	this.min = min(this.min, number)
	this.max = max(this.max, number)
	this.val[number]++
}

func (this *TwoSum) Find(value int) bool {
	if value < this.min+this.min || value > this.max+this.max {
		return false
	}
	for k, v := range this.val {
		if value-k == k {
			if v > 1 {
				return true
			}
		} else {
			if this.val[value-k] > 0 {
				return true
			}
		}
	}
	return false
}
