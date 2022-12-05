package main

import (
	"fmt"
	"strings"
)

// 71 simplify-path 2022-10-18 15:53:44

// leetcode submit region begin(Prohibit modification and deletion)
func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	stack := newStack71(len(dirs))

	for _, dir := range dirs {
		switch dir {
		case "", ".":
			continue
		case "..":
			stack.pop()
		default:
			stack.push(dir)
		}
	}

	return "/" + stack.join("/")
}

type stack71 struct {
	stack  []string
	length int
}

// 压栈
func (s *stack71) push(b string) {
	s.stack = append(s.stack, b)
	s.length++
}

func (s *stack71) pop() (res string) {
	if s.length <= 0 {
		return
	}
	s.length--
	res = s.stack[s.length]
	s.stack = s.stack[0:s.length]
	return
}

func (s *stack71) join(sep string) string {
	return strings.Join(s.stack, sep)
}

func newStack71(cap int) *stack71 {
	return &stack71{
		stack: make([]string, 0, cap),
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
}
