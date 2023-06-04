package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	response := canAttendMeetings(intervals)
	fmt.Println(response)
}

func canAttendMeetings(intervals [][]int) bool {
	alen := len(intervals)
	if alen <= 1 {
		return true
	}

	m := make(map[int]int, alen)
	s := make([]int, alen)
	for i := 0; i < alen; i++ {
		if _, ok := m[intervals[i][0]]; ok {
			return false
		}
		m[intervals[i][0]] = intervals[i][1]
		s[i] = intervals[i][0]
	}
	sort.Ints(s)
	fmt.Println(s, m)
	for j := 1; j < alen; j++ {
		if s[j] < m[s[j-1]] {
			return false
		}
	}

	return true
}
