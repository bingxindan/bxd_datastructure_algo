package main

import "fmt"

type FirstUniqChar struct {
}

func main() {
	str := "abaccdeff"
	str = "loveleetcode"
	t := new(FirstUniqChar)
	k := t.firstU(str)
	fmt.Println("k:", string(k))
}

func (t *FirstUniqChar) firstU(str string) byte {
	strs := []byte(str)
	leng := len(strs)
	tmps := make(map[byte]int, 0)

	for i := 0; i < leng; i++ {
		v := strs[i]
		tmps[v]++
	}

	for k, val := range strs {
		if tmps[val] == 1 {
			return strs[k]
		}
	}

	return ' '
}
