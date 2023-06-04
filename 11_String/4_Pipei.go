package main

import (
	"fmt"
	"strings"
)

const (
	primeRK = 16777619
)

func main() {
	s := "asdfasdfas"
	fmt.Printf("ssssss: %+v\n", s[1])
	substr := "fas"
	//idx := strings.Index(s, substr)
	idx := LastIndex(s, substr)
	fmt.Printf("aaaaaaa, s: %+s, substr: %+v, idx: %+v\n", s, substr, idx)
}

func LastIndex(s, substr string) int {
	n := len(substr)
	switch {
	case n == 0:
		return len(s)
	case n == 1:
		return strings.LastIndexByte(s, substr[0])
	case n == len(s):
		if substr == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	// Rabin-Karp search from the end of the string
	hashss := hashStrRev(substr)
	last := len(s) - n
	var h uint32
	for i := len(s) - 1; i >= last; i-- {
		h = h*primeRK + uint32(s[i])
	}
	fmt.Printf("last: %+v, h: %+v, ls: %+v, s: %+v\n", last, h, s[last:], substr)
	if h == hashss && s[last:] == substr {
		return last
	}
	return -1
}

func hashStrRev(sep string) uint32 {
	hash := uint32(0)
	for i := len(sep) - 1; i >= 0; i-- {
		hash = hash*primeRK + uint32(sep[i])
		fmt.Printf("hash: %+v, spe: %+v, u: %+v\n", hash, sep[i], uint32(sep[i]))
	}
	return hash
}
