package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat","abc","bca","aabc","baac"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"],["abc","bca"],["aabc","baac"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Example 4
Input: strs = ["ab", "ba", "ca", ""]
Output: [["ab", "ba"], ["ca"], [""]]

Constraints:
1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lower-case English letters.

N: 数组长度
M：单个字符串最大长度.
单个字符串排序复杂度：
排序算法
不确定
归并排序：M * LogM
冒泡排序：M^2
快速排序：M * LogM
选择排序：

时间复杂度: N * M * LogM
空间复杂度: N * M
*/

// 字符串排序
func sortStr(str string) string {
	strs := strings.Split(str, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}

func groupAnagrams(strs []string) [][]string {
	hashStr := make(map[string][]string, 0)
	ret := make([][]string, 0)
	leng := len(strs)

	for i := 0; i < leng; i++ {
		newStr := sortStr(strs[i])
		if v, ok := hashStr[newStr]; ok {
			hashStr[newStr] = append(v, strs[i])
		} else {
			tmp := []string{strs[i]}
			hashStr[newStr] = tmp
		}
	}

	// 循环输出最终结果
	for _, val := range hashStr {
		ret = append(ret, val)
	}

	return ret
}

func main() {
	//input := []string{"eat", "tea", "tan", "ate", "nat", "bat", ""}
	input := []string{""}
	ret := groupAnagrams(input)
	fmt.Println(ret)
}
