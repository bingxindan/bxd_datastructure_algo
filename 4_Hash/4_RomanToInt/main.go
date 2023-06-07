package main

import "fmt"

func main() {
	//s := "III"
	s := "IV"
	response := romanToInt(s)
	fmt.Println("res: ", response)
}

var symbolValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	n := len(s)
	var ans int

	for i := range s {
		fmt.Printf("i: %v, si: %v\n", i, s[i])
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}

	return ans
}
