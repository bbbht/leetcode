package main

import (
	"strings"
)

// 482 license-key-formatting 2022-12-12 10:34:23

// leetcode submit region begin(Prohibit modification and deletion)
func licenseKeyFormatting(s string, k int) string {
	dashs := strings.Count(s, "-")
	an := len(s) - dashs

	if an == 0 {
		return ""
	}

	s = strings.ToUpper(strings.Replace(s, "-", "", -1))

	res := make([]byte, (an+k-1)/k-1+an)

	i, j := len(res), len(s)
	for 0 <= j-k {
		copy(res[i-k:i], s[j-k:j])

		if 0 <= i-k-1 {
			res[i-k-1] = '-'
		}

		i -= k + 1
		j -= k
	}

	if j > 0 {
		copy(res[:j], s[:j])
	}

	return string(res)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
