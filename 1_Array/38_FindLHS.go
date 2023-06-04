package main

func main() {

}

func findLHS(nums []int) int {
	var ans int
	cnt := make(map[int]int, 0)
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if c1 := cnt[num+1]; c1 > 0 && c+c1 > ans {
			ans = c + c1
		}
	}

	return ans
}
