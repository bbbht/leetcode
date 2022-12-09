package main

// 383 ransom-note 2022-12-07 16:46:27

// leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int, len(magazine))
	for i := range magazine {
		if _, ok := m[magazine[i]]; ok {
			m[magazine[i]] += 1
		} else {
			m[magazine[i]] = 1
		}
	}

	for i := range ransomNote {
		if _, ok := m[ransomNote[i]]; !ok || m[ransomNote[i]] <= 0 {
			return false
		}
		m[ransomNote[i]] -= 1
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
