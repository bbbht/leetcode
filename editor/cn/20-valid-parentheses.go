package main

// 有效的括号 2022-07-22 18:30:12

// leetcode submit region begin(Prohibit modification and deletion)

var m20 = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	0:   0,
}

func isValid(s string) bool {
	size := len(s)
	if size%2 != 0 {
		return false
	}
	ok, s20 := false, newStack(size/2+1)
	// 左括号压入，右括号则弹出对比
	for i := range s {
		if _, ok = m20[s[i]]; ok {
			s20.push(s[i])
			// 长度超过一半，肯定对不上了
			if s20.length > size/2 {
				return false
			}
			continue
		}
		if m20[s20.pop()] != s[i] {
			return false
		}
	}

	return s20.isEmpty()
}

type stack20 struct {
	stack  []byte
	length int
}

// 压栈
func (s *stack20) push(b byte) {
	s.stack = append(s.stack, b)
	s.length++
}

func (s *stack20) pop() (res byte) {
	if s.length <= 0 {
		return
	}
	s.length--
	res = s.stack[s.length]
	s.stack = s.stack[0:s.length]
	return
}

func (s *stack20) isEmpty() bool {
	return s.length <= 0
}

func newStack(cap int) *stack20 {
	return &stack20{
		stack: make([]byte, 0, cap),
	}
}

// leetcode submit region end(Prohibit modification and deletion)
