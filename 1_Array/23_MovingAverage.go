package main

import "fmt"

func main() {
	size := 1
	obj := ConstructorII(size)
	val := 2
	response := obj.Next(val)
	fmt.Println(response)
}

type MovingAverage struct {
	size, sum int
	q         []int
}

func ConstructorII(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.q) == this.size {
		this.sum -= this.q[0]
		this.q = this.q[1:]
	}
	this.sum += val
	this.q = append(this.q, val)
	return float64(this.sum) / float64(len(this.q))
}
