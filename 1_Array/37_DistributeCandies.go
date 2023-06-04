package main

func main() {

}

func distributeCandies(candyType []int) int {
	set := make(map[int]struct{}, 0)
	for _, t := range candyType {
		set[t] = struct{}{}
	}

	ans := len(candyType) / 2
	if len(set) < ans {
		ans = len(set)
	}

	return ans
}
