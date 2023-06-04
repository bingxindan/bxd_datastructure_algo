package main

func main() {

}

func findLengthOfLCIS(nums []int) int {
	var ans int
	var start = 0

	for i, v := range nums {
		if i > 0 && v <= nums[i-1] {
			start = i
		}
		ans = max(ans, i-start+1)
	}

	return ans
}
