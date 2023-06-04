package main

import "fmt"

func main() {
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1, word2 := "makes", "coding"
	response := shortestDistance(wordsDict, word1, word2)
	fmt.Println(response)
}

func shortestDistance(wordsDict []string, word1, word2 string) int {
	ans := len(wordsDict)
	index1, index2 := -1, -1
	for i, word := range wordsDict {
		if word == word1 {
			index1 = i
		} else if word == word2 {
			index2 = i
		}
		if index1 >= 0 && index2 >= 0 {
			ans = min(ans, abs(index1-index2))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
