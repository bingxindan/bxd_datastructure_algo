package main

import "fmt"

func main() {
	s := "ltecod"
	response := isUniqueII(s)
	fmt.Println("response: ", response)
}

func isUniqueII(astr string) bool {
	var mark = 0

	for _, c := range astr {
		bit := c - 'a'
		if mark&(1<<bit) != 0 {
			return false
		}
		mark |= 1 << bit
	}

	return true
}

func isUniqueI(astr string) bool {
	n := len(astr)

	if n <= 1 {
		return true
	}

	set := map[rune]bool{}
	for _, v := range astr {
		if set[v] {
			return false
		}
		set[v] = true
	}

	return true
}

func isUnique(astr string) bool {
	n := len(astr)

	if n <= 1 {
		return true
	}

	for i := 0; i < n; i++ {
		str := astr[i]
		for j := i + 1; j <= n-1; j++ {
			str1 := astr[j]
			if str == str1 {
				return false
			}
		}
	}

	return true
}
