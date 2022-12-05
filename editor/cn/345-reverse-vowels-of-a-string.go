package main

// 345 reverse-vowels-of-a-string 2022-12-02 17:30:03

// leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	m := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	size := len(s)
	l, r, bs := 0, size-1, []byte(s)
	for {
		for l < size {
			if _, ok := m[bs[l]]; ok {
				break
			}
			l++
		}
		for r > 0 {
			if _, ok := m[bs[r]]; ok {
				break
			}
			r--
		}

		if l >= r {
			break
		}

		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}

	return string(bs)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
