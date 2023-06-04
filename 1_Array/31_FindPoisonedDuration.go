package main

import "fmt"

func main() {
	timeSeries := []int{1, 4}
	duration := 2
	response := findPoisonedDuration(timeSeries, duration)
	fmt.Println(response)
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	expired := 0
	var ans int
	for _, t := range timeSeries {
		if t >= expired {
			ans += duration
		} else {
			ans += t + duration - expired
		}
		expired = t + duration
	}
	return ans
}
