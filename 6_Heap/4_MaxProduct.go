package main

func main() {

}

func maxProduct(nums []int) int {
	a, b := nums[0], nums[1]
	if a < b {
		a, b = b, a
	}

	for _, num := range nums[2:] {
		if a < num {
			a, b = num, a
		} else if b < num {
			b = num
		}
	}

	return (a - 1) * (b - 1)
}
