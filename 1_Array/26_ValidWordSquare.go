package main

import "fmt"

func main() {
	words := []string{"ball",
		"area",
		"read",
		"lady"}
	response := validWordSquare(words)
	fmt.Println(response)
}

func validWordSquare(words []string) bool {
	for i, word := range words {
		var j int
		for _, w := range words {
			if i < len(w) {
				if j >= len(word) {
					return false
				}

				if word[j] != w[i] {
					return false
				}

				j++
			}
		}

		if j != len(word) {
			return false
		}
	}

	return true
}
