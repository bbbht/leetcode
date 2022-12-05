package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 93 restore-ip-addresses 2022-11-18 15:45:03

// leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) < 4 {
		return nil
	}

	ip, ips := make([]string, 0, 4), make([]string, 0, 4)

	dfs93(s, 0, ip, &ips)

	return ips
}

func dfs93(s string, i int, ip []string, ips *[]string) {
	size := len(s)
	if len(ip) == 4 {
		if i == size {
			*ips = append(*ips, strings.Join(ip, "."))
		}
		return
	}
	// "101023"
	for j := i; j < size; j++ {
		c := len(ip)
		// 每一段最多三位数字 0-255
		for k := 1; k <= 3 && k+j <= size; k++ {
			if !vaild93(s[j : k+j]) {
				return
			}
			ip = append(ip, s[j:k+j])
			dfs93(s, j+k, ip, ips)
			ip = ip[:len(ip)-1]
		}
		if len(ip) == c {
			return
		}
	}
}

func vaild93(s string) bool {
	if len(s) == 1 {
		return true
	}

	i, _ := strconv.Atoi(s)

	if len(s) == 2 {
		return i >= 10
	}
	if len(s) == 3 {
		return i >= 100 && i <= 255
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(restoreIpAddresses("0000"))
	// fmt.Println(restoreIpAddresses("25525511135"))
	// fmt.Println(restoreIpAddresses("101023"))
}
