package main

func main() {

}

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	max := dp[1]
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i/2] + dp[(i/2)+1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
