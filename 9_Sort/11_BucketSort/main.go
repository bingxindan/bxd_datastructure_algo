package main

import "fmt"

func main() {
	nums := []int{39, 2, 5, 23, 22, 24, 25, 26, 28}
	response := bucketSort(nums, 3)
	fmt.Println("response: ", response)
}

func bucketSort(arr []int, bucketSize int) []int {
	// 获得arr的最值
	length := len(arr)
	maxNum, minNum := arr[0], arr[0]
	for i := 0; i < length; i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
		} else if arr[i] < minNum {
			minNum = arr[i]
		}
	}
	maxValue := maxNum - minNum
	// 初始化桶
	bucketNum := maxValue/bucketSize + 1 // 桶个数
	buckets := make([][]int, bucketNum)
	for i := 0; i < bucketNum; i++ {
		buckets[i] = make([]int, 0)
	}
	// 利用映射将元素分配到每个桶中
	for i := 0; i < len(arr); i++ {
		id := (arr[i] - minNum) / bucketSize // 桶ID
		buckets[id] = append(buckets[id], arr[i])
	}
	// 对每个桶进行排序，然后按顺序将桶中数据放入arr
	arrIndex := 0
	for i := 0; i < bucketNum; i++ {
		if len(buckets[i]) == 0 { // 空桶
			continue
		}
		insertSort(buckets[i])                 // 桶内排序
		for j := 0; j < len(buckets[i]); j++ { // 将每个桶的排序结果写入arr
			arr[arrIndex] = buckets[i][j]
			arrIndex++
		}
	}
	return arr
}

func insertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ { // 无序区
		tmp := arr[i]
		left, right := 0, i-1
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] > tmp {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		j := i - 1
		for ; j >= left; j-- { // 有序区
			arr[j+1] = arr[j]
		}
		arr[left] = tmp
	}
}
