package main

func main() {

}

func getLeastNumbers(arr []int, k int) []int {
	var heapify func(nums []int, root, end int)

	heapify = func(nums []int, root, end int) {
		for {
			child := root*2 + 1
			if child > end {
				return
			}

			if child < end && nums[child] <= nums[child+1] {
				child++
			}

			if nums[root] > nums[child] {
				return
			}

			nums[root], nums[child] = nums[child], nums[root]
			root = child
		}
	}

	end := len(arr) - 1

	for i := end / 2; i >= 0; i-- {
		heapify(arr, i, end)
	}

	for i := end; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		end--
		heapify(arr, 0, end)
	}

	return arr[:k]
}
