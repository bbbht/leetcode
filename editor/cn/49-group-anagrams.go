package main

import (
	"fmt"
	"sort"
	"strings"
)

// 49 group-anagrams 2022-08-31 14:45:32

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	m := make(map[string][]string)
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		newStr := strings.Join(s, "")

		if _, ok := m[newStr]; ok {
			m[newStr] = append(m[newStr], str)
		} else {
			m[newStr] = []string{str}
		}
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
