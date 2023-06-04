package main

import "math"

func main() {

}

func findRestaurant(list1, list2 []string) []string {
	index := make(map[string]int, len(list1))
	for i, s := range list1 {
		index[s] = i
	}

	indexSum := math.MaxInt32
	ans := make([]string, 0)
	for i, s := range list2 {
		if j, ok := index[s]; ok {
			if i+j < indexSum {
				indexSum = i + j
				ans = []string{s}
			} else if i+j == indexSum {
				ans = append(ans, s)
			}
		}
	}

	return ans
}
