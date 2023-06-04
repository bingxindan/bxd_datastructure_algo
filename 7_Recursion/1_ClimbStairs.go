package main

func main() {

}

func climbStairs(n int) int {
	var tempMap = make(map[int]int)
	return climb(0, n, tempMap)
}

func climb(i, n int, tempMap map[int]int) int {
	if i == n {
		return 1
	}

	if i > n {
		return 0
	}

	if tempMap[i] == 0 {
		tempMap[i] = climb(i+1, n, tempMap) + climb(i+2, n, tempMap)
	}

	return tempMap[i]
}
