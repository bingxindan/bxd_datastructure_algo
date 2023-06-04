package main

import "fmt"

func main() {
	prices := []int{7, 6, 4, 3, 1}
	response := maxProfit(prices)
	fmt.Println(response)
}

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}

	dp := make([][]int, length)

	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for j := 1; j < length; j++ {
		dp[j][0] = max(dp[j-1][0], -prices[j])
		dp[j][1] = max(dp[j-1][1], dp[j-1][0]+prices[j])
	}

	return dp[length-1][1]
}
