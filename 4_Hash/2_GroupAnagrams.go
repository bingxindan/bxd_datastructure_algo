package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string, 0)

	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	ans := make([][]string, 0, len(mp))

	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}
