package main

// 1598 crawler-log-folder 2022-09-09 10:49:13

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(logs []string) int {
	dirs := make([]struct{}, 0, len(logs)/2)
	for i := range logs {
		switch logs[i] {
		case "./":
		case "../":
			if len(dirs) > 0 {
				dirs = dirs[:len(dirs)-1]
			}
		default:
			dirs = append(dirs, struct{}{})
		}
	}
	return len(dirs)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
