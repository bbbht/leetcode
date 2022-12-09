package main

// 389 find-the-difference 2022-12-07 17:06:49

// leetcode submit region begin(Prohibit modification and deletion)
func findTheDifference(s string, t string) byte {
	var res byte
	// 与前面一道题一样的逻辑
	// 只要两个字符串合并，那么就是 2n+1 结果就是那个 1
	for i := range s {
		res ^= s[i]
	}
	for i := range t {
		res ^= t[i]
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
